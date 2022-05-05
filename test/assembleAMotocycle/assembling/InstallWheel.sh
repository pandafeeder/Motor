#!/bin/bash
MRequireInput data/FrameWithForkWithEngineWithExhaustWithGasTank
MRequireInput data/PumpedUpWheel
MGenOutput data/FrameWithForkWithEngineWithExhaustWithGasTankWithPumpedUpWheel

echo "Install PumpedUpWhell on Frame"
sleep 11
touch data/FrameWithForkWithEngineWithExhaustWithGasTankWithPumpedUpWheel
echo "Done"

exit

