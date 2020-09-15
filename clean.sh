echo Initializing HLF-IOTA Connector cleaning...

# Change to home directory
cd

# Stop and clean HLF 
cd /home/alejandro/connector/hyperledger-fabric-demo-app/basic-network
./stop.sh
./teardown.sh

# Kill docker ps and connector docker image
echo Removing docker ps and connector image
cd ..
python3 remove-connector-image.py

# Remove keys 
echo Removing keys
cd ~ && rm -rf .hfc-key-store/

echo Done!
