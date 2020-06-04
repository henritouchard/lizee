

BUILTFOLDER=./frontBuild

echo "run install && build react app"
cd ./app && npm install && npm run build
cd ../
if [ -d "$BUILTFOLDER" ]; then
    sudo docker-compose up --b
else
    echo "ERROR: $BUILTFOLDER does not exist. An error probably occured during npm build. You can build it by yourself in ./app folder"
    exit 1
fi

