#!/bin/bash
MRequireInput data/Wheel
MGenOutput data/PumpedUpWheel

echo "Pumpping up the whell"
sleep 5
touch data/PumpedUpWheel
echo "Done"

exit
