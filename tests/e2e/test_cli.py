import re

import pexpect
import pytest
import getpass
from helpers import run, run_interaction, \
    TEST_NET_NETWORK, TEST_NET_NODE_TCP, TEST_NET_NODE_HTTP, TEST_NET_DESTINATION, TEST_NET_DESTINATION_HTTP, TEST_NET_FEES, TEST_NET_GAS_X_GAS_PRICES, YES_FLAG, \
    SENDER_ADDRESS, RECEIVER_ADDRESS, CODE_0


@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("help", "",r"cheqd App(.*?)Usage:(.*?)Available Commands:(.*?)Flags:"),
            ("status", TEST_NET_NODE_TCP, fr"\"NodeInfo\"(.*?)\"network\":\"{TEST_NET_NETWORK}\"(.*?)\"moniker\":\"seed1-us-testnet-cheqd\""),
            ("status", TEST_NET_NODE_HTTP, fr"\"NodeInfo\"(.*?)\"network\":\"{TEST_NET_NETWORK}\"(.*?)\"moniker\":\"node1-eu-testnet-cheqd\""),
        ]
    )
def test_basic(command, params, expected_output):
    command_base = "cheqd-noded"
    run(command_base, command, params, expected_output)


@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("add", "test1", r"- name: test1(.*?)type: local(.*?)address: (.*?)pubkey: (.*?)mnemonic: "),
            ("list", None, "- name: test1"),
            ("delete", f"test1 {YES_FLAG}", r"Key deleted forever \(uh oh!\)"),
            ("add", "test2", "- name: test2"),
            ("show", "test2", "- name: test2"),
            ("delete", f"test2 {YES_FLAG}", r"Key deleted forever \(uh oh!\)"),
            ("show", "test9", "Error: test9 is not a valid name or address"),
            ("mnemonic", None, '(\w+\s){23}(\w+){1}\r\n')
        ]
    )
def test_keys(command, params, expected_output):
    command_base = "cheqd-noded keys"
    run(command_base, command, params, expected_output)


@pytest.fixture(scope='session')
def create_export_keys():
    command_base = "cheqd-noded keys"
    run(command_base, "add", "export_key", "name: export_key")


# tbd - import, migrate, parse
@pytest.mark.usefixtures('create_export_keys')
@pytest.mark.parametrize(
    "command, params, expected_output, input_string, expected_output_2",
    [
        ("export", "export_key", "Enter passphrase to encrypt the exported key", "123456", "password must be at least 8 characters"),
        ("export", "export_key", "Enter passphrase to encrypt the exported key", "12345678",
        "BEGIN TENDERMINT PRIVATE KEY"),
        ("export", "export_key", "Enter passphrase to encrypt the exported key", "qwe!@#$%^123",
         "BEGIN TENDERMINT PRIVATE KEY"),
        ("export", "export_key", "Enter passphrase to encrypt the exported key", "ttttttttttttttttttttttttttttttttttttttt",
         "BEGIN TENDERMINT PRIVATE KEY"),
    ]
)
def test_keys_export(command, params, expected_output, input_string, expected_output_2):
    command_base = "cheqd-noded keys"
    cli = run(command_base, command, params, expected_output)
    run_interaction(cli, input_string, expected_output_2)


@pytest.fixture(scope='session')
def create_import_keys():
    command_base = "cheqd-noded keys"
    run(command_base, "add", "import_key", "name: import_key")


@pytest.mark.usefixtures('create_import_keys')
def test_keys_import():
    command_base = "cheqd-noded keys"
    key_name = "import_key"
    passphrase = "00000000"
    filename = "/home/{}/key.txt".format(getpass.getuser())
    cli = run(command_base, "export", key_name, "Enter passphrase")
    run_interaction(cli, passphrase, "BEGIN")

    key_content = "-----BEGIN{}".format(cli.read())

    text_file = open(filename, "w")
    text_file.write(key_content)
    text_file.close()

    cli = run(command_base, "delete", key_name, "Continue?")
    run_interaction(cli, "y", "Key deleted forever")

    cli = run(command_base, "import", "{} {}".format(key_name, filename), "Enter passphrase to decrypt your key")
    run_interaction(cli, passphrase, "")

    run(command_base, "list", None, "- name: {}".format(key_name))


@pytest.mark.skip
@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("staking validators", f"{TEST_NET_DESTINATION}", r"pagination:(.*?)validators:"),
            ("bank balances", f"{SENDER_ADDRESS} {TEST_NET_DESTINATION}", r"balances:(.*?)amount:(.*?)denom: ncheq(.*?)pagination:"),
        ]
    )
def test_query(command, params, expected_output):
    command_base = "cheqd-noded query"
    run(command_base, command, params, expected_output)


@pytest.mark.usefixtures('restore_test_keys') # for pipeline
@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("bank send", "", r"Error: accepts 3 arg\(s\), received 0"), # no args
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 0ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"Error: : invalid coins"), # 0
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 100ncheq {TEST_NET_FEES} {YES_FLAG}", r"Error: post failed: Post \"http://localhost:26657\"(.*?)connect: connection refused"), # no destination, localhost error
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 1ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"1ncheq\""), # 1 + fees
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 2ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"2ncheq\""), # 2 + gas x price
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 99ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"99ncheq\""),
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 1ncheq {TEST_NET_DESTINATION} {YES_FLAG}", r"\"code\":13(.*?)insufficient fees"), # no fees
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 2ncheq {TEST_NET_DESTINATION} --fees 4000000ncheq {YES_FLAG}", r"\"code\":13(.*?)insufficient fees"), # bad fees
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 3ncheq {TEST_NET_DESTINATION} --gas 70000 --gas-prices 1ncheq {YES_FLAG}", r"\"code\":13(.*?)insufficient fees"), # bad gas price
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 4ncheq {TEST_NET_DESTINATION} --gas 1 --gas-prices 25ncheq {YES_FLAG}", r"\"code\":11(.*?)out of gas"), # bad gas amount
            ("bank send", f"{RECEIVER_ADDRESS} {SENDER_ADDRESS} 2ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"2ncheq\""), # transfer back: 2 + fees
            ("bank send", f"{RECEIVER_ADDRESS} {SENDER_ADDRESS} 1ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"1ncheq\""), # transfer back: 1 + gas x price
            ("bank send", f"{RECEIVER_ADDRESS} {SENDER_ADDRESS} 999999999ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"\"code\":5(.*?)insufficient funds"),
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 1000ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG} --note 'test123!=$'", fr"{CODE_0}(.*?)\"value\":\"1000ncheq\""),
            ("bank send", f"{SENDER_ADDRESS} {RECEIVER_ADDRESS} 9999ncheq {TEST_NET_DESTINATION_HTTP} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"9999ncheq\""), # http + gas x price
            ("bank send", f"{RECEIVER_ADDRESS} {SENDER_ADDRESS} 9999ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"9999ncheq\""), # transfer back: tcp + fees
        ]
    )
def test_tx(command, params, expected_output):
    command_base = "cheqd-noded tx"
    run(command_base, command, params, expected_output)


@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("show-address", "", r"cheqd(.*?)"),
            ("show-node-id", "", r"^[a-z\d]{40}"),
            ("show-validator", "", r"\"\@type\":(.*?)\"key\":"),
        ]
    )
def test_tendermint(command, params, expected_output):
    command_base = "cheqd-noded tendermint"
    run(command_base, command, params, expected_output)


def test_production(send_with_note):
    tx_hash, tx_memo = send_with_note
    run("cheqd-noded query", "tx", f"{tx_hash} {TEST_NET_DESTINATION}", fr"code: 0(.*?)memo: {tx_memo}(.*?)txhash: {tx_hash}") # check that txn has correct memo value
