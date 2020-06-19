gen:
	minimock -g -i Client -o ./ -s _mock.go
	easyjson api.go notification.go
