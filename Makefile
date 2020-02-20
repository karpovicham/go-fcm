gen:
	easyjson -all notification.go response.go
	minimock -g -i Client -o ./ -s _mock.go