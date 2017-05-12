# Start SugarChar Application
cd ./store && go get && go install &&
cd ../handlers && go get && go install &&
cd ../psql && go get && go install &&
cd ../ && go get && go install &&
sugarchat
