sh build_wasm.sh;

{
     sleep 1 
     echo  "\n"
     echo "    Game Server:    http://localhost:9080"
}&

# Then open the "http://localhost:9080" on the web
go run main.go
