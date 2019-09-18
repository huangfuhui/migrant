// Copyright 2019 Migrant
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(driver, DNS string) (err error) {
	DB, err = sql.Open(driver, DNS)
	if err != nil {
		return
	}
	err = DB.Ping()
	return
}

// When all is done, it is a good habit to close the database.
func CloseDB() error {
	return DB.Close()
}
