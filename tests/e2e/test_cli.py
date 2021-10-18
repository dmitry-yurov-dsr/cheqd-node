import sys
import os
import re
import pexpect
import pytest


IMPLICIT_TIMEOUT = 30
ENCODING = "utf-8"
READ_BUFFER = 6000
TEST_NET_DESTINATION = "--node 'tcp://seed1.us.testnet.cheqd.network:26657' --chain-id 'cheqd-testnet-2'"
TEST_NET_FEES = "--fees 5000000ncheq"
TEST_NET_GAS_X_GAS_PRICES = "--gas 70000 --gas-prices 25ncheq"
YES_FLAG = "-y"

sender = "cheqd1ece09txhq6nm9fkft9jh3mce6e48ftescs5jsw"
receiver = "cheqd16d72a6kusmzml5mjhzjv63c9j5xnpsyqs8f3sk"


def run(command_base, command, params, expected_output):
    cli = pexpect.spawn(f"{command_base} {command} {params}", encoding=ENCODING, timeout=IMPLICIT_TIMEOUT, maxread=READ_BUFFER)
    cli.logfile = sys.stdout
    cli.expect(expected_output)
    print(f"BEFORE >>> {cli.before}")
    print(f"AFTER >>> {cli.after}")
    return cli


@pytest.mark.skip
@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("help", "",r"cheqd App(.*?)Usage:(.*?)Available Commands:(.*?)Flags:"),
            # ("version", "",os.environ["RELEASE_NUMBER"]), # this works against deb package but not against starport build
            ("status", "",r"\"NodeInfo\"(.*?)\"network\":\"cheqd\"(.*?)\"moniker\":\"node0\""),
        ]
    )
def test_basic(command, params, expected_output):
    command_base = "cheqd-noded"
    run(command_base, command, params, expected_output)


@pytest.mark.skip
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
        ]
    )
def test_keys(command, params, expected_output):
    command_base = "cheqd-noded keys"
    run(command_base, command, params, expected_output)


# use fixture to send with note and then check memo
@pytest.mark.skip
@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("staking validators", f"{TEST_NET_DESTINATION}", r"pagination:(.*?)validators:"),
            ("bank balances", f"{sender} {TEST_NET_DESTINATION}", r"balances:(.*?)amount:(.*?)denom: ncheq(.*?)pagination:"),
            # ("tx", f"{tx_hash} {TEST_NET_DESTINATION}", f"{tx_memo}"),
        ]
    )
def test_query(command, params, expected_output):
    command_base = "cheqd-noded query"
    run(command_base, command, params, expected_output)


# use fixture to recover sender + receiver keys
@pytest.mark.skip
@pytest.mark.parametrize(
        "command, params, expected_output",
        [
            ("bank send", "", r"Error: accepts 3 arg\(s\), received 0"), # no args
            ("bank send", f"{sender} {receiver} 0ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"Error: : invalid coins"), # 0
            ("bank send", f"{sender} {receiver} 1ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"\"code\":0(.*?)\"value\":\"1ncheq\""), # 1 + fees
            ("bank send", f"{sender} {receiver} 2ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", r"\"code\":0(.*?)\"value\":\"2ncheq\""), # 2 + gas x price
            ("bank send", f"{sender} {receiver} 99ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"\"code\":0(.*?)\"value\":\"99ncheq\""),
            ("bank send", f"{sender} {receiver} 99ncheq {TEST_NET_DESTINATION} {YES_FLAG}", r"\"code\":13(.*?)insufficient fees"),
            ("bank send", f"{receiver} {sender} 2ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"\"code\":0(.*?)\"value\":\"2ncheq\""), # transfer back 2 + fees
            ("bank send", f"{receiver} {sender} 1ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", r"\"code\":0(.*?)\"value\":\"1ncheq\""), # transfer back 1 + gas x price
            ("bank send", f"{receiver} {sender} 999999999ncheq {TEST_NET_DESTINATION} {TEST_NET_FEES} {YES_FLAG}", r"\"code\":5(.*?)insufficient funds"),
            ("bank send", f"{sender} {receiver} 1000ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG} --note 'test123!=$'", r"\"code\":0(.*?)\"value\":\"1000ncheq\""),
        ]
    )
def test_tx(command, params, expected_output):
    command_base = "cheqd-noded tx"
    run(command_base, command, params, expected_output)


# @pytest.mark.skip
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


@pytest.mark.skip
def test_production():
    note = "test_memo"
    cli = run("cheqd-noded tx", "bank send", f"{sender} {receiver} 1000ncheq {TEST_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG} --note {note}", r"\"code\":0(.*?)\"value\":\"1000ncheq\"")
    tx_hash = re.search(r"\"txhash\":\"(.+?)\"", cli.before).group(1).strip()
    print(tx_hash)
