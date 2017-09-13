// gocouch project main.go
// register main handlers for the Go CouchDB web server
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"GConfig"
	"Internals"
	"JSONStructs"
)

var (
	SlashRSP     *JSONStructs.SlashResponse
	ValidDB_Name *regexp.Regexp
	e            *echo.Echo

	//DBList map[string]Internals.BoltDB
)

func slashHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	c.JSON(http.StatusOK, SlashRSP)
	return nil
}

func uuidHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("ETag", "buru-buru")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	defUUIDS := []string{Internals.GetUUID(),
		Internals.GetUUID(),
		Internals.GetUUID(),
		Internals.GetUUID(),
		Internals.GetUUID()}

	UUIDRSP := &JSONStructs.UUIDSResponse{
		UUIDS: defUUIDS,
	}

	c.JSON(http.StatusOK, UUIDRSP)

	return nil
}

func dbHeadHandler(c echo.Context) error {
	db := c.Param("db")

	if Contains(ALL_DBS, db) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
		c.Response().Header().Set("Cache-Control", "must-revalidate")
		c.Response().Header().Set("Server", Internals.ServerMsg)
	} else {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
		c.Response().Header().Set("Cache-Control", "must-revalidate")
		c.Response().Header().Set("Server", Internals.ServerMsg)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	//json.NewEncoder(c.Response()).Encode(c.Param("db"))
	c.Response().Flush()
	//c.JSON(http.StatusOK, rr)
	return nil
}

func dbPutHandler(c echo.Context) error {
	db := c.Param("db")

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	if !ValidDB_Name.MatchString(db) {
		//the DB name is not valid -> 400

		InvalidDBName := &JSONStructs.DBErrorResponse{
			ErrorMsg: "illegal_database_name",
			Reason: "Name: '" + db + "'. Only lowercase characters (a-z), " +
				"digits (0-9), and any of the characters _, $, " +
				"(, ), +, -, and / are allowed. Must begin with " +
				" a letter.",
		}
		c.JSON(http.StatusBadRequest, InvalidDBName)
		return nil
	}
	if Contains(ALL_DBS, db) {
		//the DB already exists -> 412
		DBExists := &JSONStructs.DBErrorResponse{
			ErrorMsg: "file_exists",
			Reason:   "The database could not be created, the file already exists.",
		}
		c.JSON(http.StatusPreconditionFailed, DBExists)
		return nil
	}
	//TODO implement database security so that 401 may be checked also

	//It should be OK - create the DB file and add it to ALL_DBS
	//Send 201 response

	newDB := Internals.NewBoltDB("", db+".bd")
	defer newDB.Close()
	ALL_DBS = append(ALL_DBS, db)

	c.Response().Header().Set("Location", "http://get.server.address.or.host.name"+":"+
		GConfig.GoCouchCFG.HTTPd.Port+"/"+db)
	DBCreatedOK := &JSONStructs.OKResponse{
		OK: true,
	}
	c.JSON(http.StatusCreated, DBCreatedOK)
	return nil
}

func dbPostHandler(c echo.Context) error {
	db := c.Param("db")

	//check db
	//400 or 404 possible error

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	if !ValidDB_Name.MatchString(db) {
		//the DB name is not valid -> 400

		InvalidDBName := &JSONStructs.DBErrorResponse{
			ErrorMsg: "illegal_database_name",
			Reason: "Name: '" + db + "'. Only lowercase characters (a-z), " +
				"digits (0-9), and any of the characters _, $, " +
				"(, ), +, -, and / are allowed. Must begin with " +
				" a letter.",
		}
		c.JSON(http.StatusBadRequest, InvalidDBName)
		return nil
	}
	if !Contains(ALL_DBS, db) {
		//the DB does not exist -> 404
		DBExists := &JSONStructs.DBErrorResponse{
			ErrorMsg: "not_found",
			Reason:   "no_db_file",
		}
		c.JSON(http.StatusNotFound, DBExists)
		return nil
	}
	//TODO implement database security so that 401 may be checked also
	//implement batch transactions and return 202 Accepted

	//Get JSON payload and decode it
	dec := json.NewDecoder(c.Request().Body)
	var payload map[string]interface{}
	for {

		if err := dec.Decode(&payload); err == io.EOF {
			break
		} else if err != nil {
			e.Logger.Error("%s", err)
		}
		//e.Logger().Info("%s", payload["_id"])
	}

	if payload["_id"] != nil {
		//check to see if there is already a document with this id -> 409
		tmp_db := Internals.NewBoltDB("", db+".bd")
		defer tmp_db.Close()
		if tmp_db.ExistsDoc([]byte(payload["_id"].(string))) {
			DOCExists := &JSONStructs.DBErrorResponse{
				ErrorMsg: "conflict",
				Reason:   "Document update conflict.",
			}
			c.JSON(http.StatusConflict, DOCExists)
			return nil
		} else {
			//create new doc with specified id
			doc, _ := json.Marshal(payload)
			payload["_rev"] = "1-" + Internals.GetMD5Hash(doc)
			doc, _ = json.Marshal(payload)

			//e.Logger().Info("%s", doc)

			tmp_db.UpdateDB([]byte(payload["_id"].(string)),
				[]byte(payload["_rev"].(string)),
				[]byte(doc))
		}
	} else {
		//create a new doc with random id
		tmp_db := Internals.NewBoltDB("", db+".bd")
		defer tmp_db.Close()
		payload["_id"] = Internals.GetUUID()
		doc, _ := json.Marshal(payload)
		payload["_rev"] = "1-" + Internals.GetMD5Hash(doc)
		doc, _ = json.Marshal(payload)

		//e.Logger().Info("%s", doc)

		tmp_db.UpdateDB([]byte(payload["_id"].(string)),
			[]byte(payload["_rev"].(string)),
			[]byte(doc))
	}

	DOCCreated := &JSONStructs.DocOKResponse{
		OK:  true,
		ID:  payload["_id"].(string),
		REV: payload["_rev"].(string),
	}

	c.Response().Header().Set("ETag", payload["_rev"].(string))
	c.Response().Header().Set("Location", "http://get.server.address.or.host.name"+":"+
		GConfig.GoCouchCFG.HTTPd.Port+"/"+db+"/"+
		payload["_id"].(string))
	c.JSON(http.StatusCreated, DOCCreated)
	return nil
}

func dbDeleteHandler(c echo.Context) error {
	db := c.Param("db")
	rev := c.QueryParam("rev")

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Server", Internals.ServerMsg)
	//Check for possible errors 400 and 404
	if !Contains(ALL_DBS, db) {
		//the DB does not exist -> 404
		DBExists := &JSONStructs.DBErrorResponse{
			ErrorMsg: "not_found",
			Reason:   "missing",
		}
		c.JSON(http.StatusNotFound, DBExists)
		return nil
	}

	if len(rev) > 0 || !ValidDB_Name.MatchString(db) {
		//the DB name is not valid -> 400

		InvalidDBName := &JSONStructs.DBErrorResponse{
			ErrorMsg: "bad_reques",
			Reason: "You tried to DELETE a database with a ?rev= parameter. " +
				"Did you mean to DELETE a document instead?",
		}
		c.JSON(http.StatusBadRequest, InvalidDBName)
		return nil
	}

	//TODO implement database security so that 401 may be checked also
	err := Internals.DeleteDB("", db+".bd")
	if err != nil {
		e.Logger.Error("%s", err)
	}
	//Remove also from ALL_DBS
	all_dbs_init("")
	DeleteOKResponse := &JSONStructs.OKResponse{
		OK: true,
	}
	c.JSON(http.StatusOK, DeleteOKResponse)
	return nil
}

func dbBackup(c echo.Context) error {
	db := c.Param("db")

	//Check the existence of the database -> 404
	c.Response().Header().Set("Server", Internals.ServerMsg)
	//Check for possible errors 400 and 404
	if !Contains(ALL_DBS, db) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().Header().Set("Cache-Control", "must-revalidate")
		//the DB does not exist -> 404
		DBExists := &JSONStructs.DBErrorResponse{
			ErrorMsg: "not_found",
			Reason:   "Missing database file!",
		}
		c.JSON(http.StatusNotFound, DBExists)
		return nil
	}

	tmp_db := Internals.ROBoltDB("", db+".bd")
	defer tmp_db.Close()
	err := tmp_db.ExportFile("", db+".bd", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Flush()
	return nil
}

func dbGetHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("ETag", "buru-buru")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	json.NewEncoder(c.Response()).Encode(c.Param("db"))
	c.Response().Flush()
	//c.JSON(http.StatusOK, rr)
	return nil
}

func all_dbsHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("Server", Internals.ServerMsg)

	c.JSON(http.StatusOK, ALL_DBS)
	return nil
}

func init() {
	ValidDB_Name = regexp.MustCompile(`^[a-z][a-z0-9_$()+/-]*$`)
	GConfig.CheckCFGFile()
	GConfig.CheckReplicatorFile()
	all_dbs_init("")

	SlashRSP = &JSONStructs.SlashResponse{
		CouchDB: "Welcome to GO CouchDB",
		UUID:    GConfig.GoCouchCFG.GoCouch.UUID,
		Version: "0.0.1",
		Vendor:  map[string]string{"version": "0.0.1", "vendor": "Free Time Software"},
	}

}

func main() {

	e = echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	//Register routes
	e.GET("/", slashHandler)
	e.GET("/_uuids", uuidHandler)
	e.GET("/_all_dbs", all_dbsHandler)

	//Database operations
	e.HEAD("/:db", dbHeadHandler)
	e.GET("/:db", dbGetHandler)
	e.PUT("/:db", dbPutHandler)
	e.POST("/:db", dbPostHandler)
	e.DELETE("/:db", dbDeleteHandler)
	//_backup
	e.GET("/_backup/:db", dbBackup)

	e.Logger.Info("%s", Internals.WelcomeMsg)

	//Start server
	e.Logger.Fatal(e.Start(GConfig.GoCouchCFG.HTTPd.Bind_Address +
		":" + GConfig.GoCouchCFG.HTTPd.Port))
}
