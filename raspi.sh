echo "Building ..."
env GOOS=linux GOARCH=arm go build -o dist/linux-arm .
echo "Uploading ..."
scp dist/linux-arm admin@bigbro.local:/home/admin
