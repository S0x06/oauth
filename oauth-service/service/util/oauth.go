package util

import (
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func ManagerOauth(appid string, cf &models.Client) server ,error{

	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store

	clientStore := store.NewClientStore()

//	&models.Client{
//		ID:     "000000",
//		Secret: "999999",
//		Domain: "http://localhost"}

	clientStore.Set(appid, cf)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return srv, err.Error()
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		return srv,re.Error.Error()
	})
	
	return srv, nil

}
