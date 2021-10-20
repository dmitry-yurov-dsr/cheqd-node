import sys
import os
import pexpect
import re
import random
import string


IMPLICIT_TIMEOUT = 30
ENCODING = "utf-8"
READ_BUFFER = 6000

TEST_NET_NETWORK = "cheqd-testnet-2"
LOCAL_NET_NETWORK = "cheqd"
TEST_NET_NODE_TCP = "--node 'tcp://seed1.us.testnet.cheqd.network:26657'"
TEST_NET_NODE_HTTP = "--node http://node1.eu.testnet.cheqd.network:26657/"
LOCAL_NET_NODE_TCP = "--node 'tcp://localhost:26657'"
LOCAL_NET_NODE_HTTP = "--node http://localhost:26657/"
TEST_NET_DESTINATION = f"{TEST_NET_NODE_TCP} --chain-id 'cheqd-testnet-2'"
TEST_NET_DESTINATION_HTTP = f"{TEST_NET_NODE_HTTP} --chain-id 'cheqd-testnet-2'"
LOCAL_NET_DESTINATION = f"{LOCAL_NET_NODE_TCP} --chain-id 'cheqd'"
LOCAL_NET_DESTINATION_HTTP = f"{LOCAL_NET_NODE_HTTP} --chain-id 'cheqd'"
TEST_NET_FEES = "--fees 5000000ncheq"
TEST_NET_GAS_X_GAS_PRICES = "--gas 70000 --gas-prices 25ncheq"
TEST_NET_GAS_X_GAS_PRICES_INT = 1750000
YES_FLAG = "-y"

SENDER_ADDRESS = "cheqd1ece09txhq6nm9fkft9jh3mce6e48ftescs5jsw"
SENDER_MNEMONIC = "oil long siege student rent jar awkward park entry ripple enable company sort people little damp arrange wise slender push brief solve tattoo cycle"
RECEIVER_ADDRESS= "cheqd16d72a6kusmzml5mjhzjv63c9j5xnpsyqs8f3sk"
RECEIVER_MNEMONIC = "strike impact earth indoor man illness virus genuine rib control antenna loop neck rotate bargain original nasty size either try snap quiz stairs huge"

LOCAL_SENDER_ADDRESS = os.environ["OP0_ADDRESS"]
LOCAL_RECEIVER_ADDRESS = os.environ["OP1_ADDRESS"]

CODE_0 = "\"code\":0"


def random_string(length):
    return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))


def run(command_base, command, params, expected_output):
    cli = pexpect.spawn(f"{command_base} {command} {params}", encoding=ENCODING, timeout=IMPLICIT_TIMEOUT, maxread=READ_BUFFER)
    cli.logfile = sys.stdout
    cli.expect(expected_output)
    print(f"BEFORE >>> {cli.before}") # FIXME DEBUG
    print(f"AFTER >>> {cli.after}") # FIXME DEBUG
    return cli


def run_interaction(cli, input_string, expected_output):
    cli.sendline(input_string)
    cli.expect(expected_output)


def get_balance(address, network_destination):
    cli = run("cheqd-noded query", "bank balances", f"{address} {network_destination}", r"balances:(.*?)amount:(.*?)denom: ncheq(.*?)pagination:")
    balance = re.search(r"amount: \"(.+?)\"", cli.after).group(1).strip()
    return balance


def send_with_note(note):
    cli = run("cheqd-noded tx", "bank send", f"{LOCAL_SENDER_ADDRESS} {LOCAL_RECEIVER_ADDRESS} 1000ncheq {LOCAL_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG} --note {note}", fr"{CODE_0}(.*?)\"value\":\"1000ncheq\"")
    tx_hash = re.search(r"\"txhash\":\"(.+?)\"", cli.before).group(1).strip()
    return tx_hash, note


def set_up_operator():
    name = random_string(10)
    cli = run("cheqd-noded keys", "add", name, r"mnemonic: \"\"")
    address = re.search(r"address: (.+?)\n", cli.before).group(1).strip()
    print(address)
    pubkey = re.search(r"pubkey: (.+?)\n", cli.before).group(1).strip()
    print(pubkey)
    run("cheqd-noded tx", "bank send", f"{LOCAL_SENDER_ADDRESS} {address} 1100000000000000ncheq {LOCAL_NET_DESTINATION} {TEST_NET_GAS_X_GAS_PRICES} {YES_FLAG}", fr"{CODE_0}(.*?)\"value\":\"1100000000000000ncheq\"")
    return name, address, pubkey