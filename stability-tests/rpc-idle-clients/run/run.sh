#!/bin/bash
rm -rf /tmp/pugdagd-temp

NUM_CLIENTS=128
pugdagd --devnet --appdir=/tmp/pugdagd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
KASPAD_PID=$!
KASPAD_KILLED=0
function killPugdagdIfNotKilled() {
  if [ $KASPAD_KILLED -eq 0 ]; then
    kill $KASPAD_PID
  fi
}
trap "killPugdagdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $KASPAD_PID

wait $KASPAD_PID
KASPAD_EXIT_CODE=$?
KASPAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Pugdagd exit code: $KASPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
