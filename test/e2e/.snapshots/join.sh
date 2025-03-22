#!/bin/bash

rm -rf test/e2e/snapshots
mkdir -p test/e2e/snapshots
pushdir test/e2e/snapshots
gh run download 14004127672 -p "snapshots_*"

# loop all files
for dir in snapshots_*/
do
    cd $dir
    mv * ..
    cd ..
    rm -rf $dir
done
popdir
