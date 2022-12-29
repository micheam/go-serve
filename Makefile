TARGET = serve
PHONY : clean install

$(TARGET) : main.go
	go build -o bin/$(TARGET) .

clean :
	rm -f bin/$(TARGET)

install : clean $(TARGET)
	cp bin/$(TARGET) ${HOME}/bin/

