// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1078
	`ALTER`: {
		//line sql.y: 1079
		Category: hGroup,
		//line sql.y: 1080
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1094
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1095
		Category: hDDL,
		//line sql.y: 1096
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1131
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1145
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1146
		Category: hDDL,
		//line sql.y: 1147
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1149
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1156
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1157
		Category: hDDL,
		//line sql.y: 1158
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1181
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1182
		Category: hPriv,
		//line sql.y: 1183
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1185
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1190
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1191
		Category: hDDL,
		//line sql.y: 1192
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1194
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1202
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1203
		Category: hDDL,
		//line sql.y: 1204
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1215
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1220
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1221
		Category: hDDL,
		//line sql.y: 1222
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1230
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1604
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1605
		Category: hCCL,
		//line sql.y: 1606
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1623
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1631
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1632
		Category: hCCL,
		//line sql.y: 1633
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1649
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1667
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1668
		Category: hCCL,
		//line sql.y: 1669
		Text: `
IMPORT [ TABLE <tablename> FROM ]
       <format> ( <datafile> )
       [ WITH <option> [= <value>] [, ...] ]

IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1695
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1724
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1725
		Category: hCCL,
		//line sql.y: 1726
		Text: `
EXPORT INTO <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1735
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1825
	`CANCEL`: {
		//line sql.y: 1826
		Category: hGroup,
		//line sql.y: 1827
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1834
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1835
		Category: hMisc,
		//line sql.y: 1836
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1839
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1857
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1858
		Category: hMisc,
		//line sql.y: 1859
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1862
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1893
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1894
		Category: hMisc,
		//line sql.y: 1895
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1898
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1947
	`CREATE`: {
		//line sql.y: 1948
		Category: hGroup,
		//line sql.y: 1949
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2028
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 2029
		Category: hExperimental,
		//line sql.y: 2030
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 2088
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2089
		Category: hDML,
		//line sql.y: 2090
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2094
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2109
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2110
		Category: hCfg,
		//line sql.y: 2111
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2123
	`DROP`: {
		//line sql.y: 2124
		Category: hGroup,
		//line sql.y: 2125
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2142
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2143
		Category: hDDL,
		//line sql.y: 2144
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2145
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2157
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2158
		Category: hDDL,
		//line sql.y: 2159
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2160
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2172
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2173
		Category: hDDL,
		//line sql.y: 2174
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2175
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2187
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2188
		Category: hDDL,
		//line sql.y: 2189
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2190
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2210
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2211
		Category: hDDL,
		//line sql.y: 2212
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2213
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2233
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2234
		Category: hPriv,
		//line sql.y: 2235
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2236
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2248
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2249
		Category: hPriv,
		//line sql.y: 2250
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2251
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2273
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2274
		Category: hMisc,
		//line sql.y: 2275
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2288
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2348
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2349
		Category: hMisc,
		//line sql.y: 2350
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2351
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2373
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2374
		Category: hMisc,
		//line sql.y: 2375
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2376
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2397
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2398
		Category: hMisc,
		//line sql.y: 2399
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2400
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2420
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2421
		Category: hPriv,
		//line sql.y: 2422
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2435
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2451
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2452
		Category: hPriv,
		//line sql.y: 2453
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2466
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2521
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2522
		Category: hCfg,
		//line sql.y: 2523
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2524
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2536
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2537
		Category: hCfg,
		//line sql.y: 2538
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2539
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2548
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2549
		Category: hCfg,
		//line sql.y: 2550
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2553
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2574
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2575
		Category: hExperimental,
		//line sql.y: 2576
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2584
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2590
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2591
		Category: hExperimental,
		//line sql.y: 2592
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2600
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2608
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2609
		Category: hExperimental,
		//line sql.y: 2610
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2621
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2676
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2677
		Category: hCfg,
		//line sql.y: 2678
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2679
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2700
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2701
		Category: hCfg,
		//line sql.y: 2702
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2708
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2725
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2726
		Category: hTxn,
		//line sql.y: 2727
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2734
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2902
	`SHOW`: {
		//line sql.y: 2903
		Category: hGroup,
		//line sql.y: 2904
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW JOBS,
SHOW QUERIES, SHOW ROLES, SHOW SESSION, SHOW SESSIONS, SHOW STATISTICS,
SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 2936
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2937
		Category: hCfg,
		//line sql.y: 2938
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2939
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2960
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 2961
		Category: hExperimental,
		//line sql.y: 2962
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2969
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2983
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 2984
		Category: hExperimental,
		//line sql.y: 2985
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2989
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3003
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3004
		Category: hCCL,
		//line sql.y: 3005
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3006
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3033
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3034
		Category: hCfg,
		//line sql.y: 3035
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3038
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3055
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3056
		Category: hDDL,
		//line sql.y: 3057
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3058
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3066
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3067
		Category: hDDL,
		//line sql.y: 3068
		Text: `SHOW DATABASES
`,
		//line sql.y: 3069
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3077
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3078
		Category: hPriv,
		//line sql.y: 3079
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3085
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3098
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3099
		Category: hDDL,
		//line sql.y: 3100
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3101
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3119
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3120
		Category: hDDL,
		//line sql.y: 3121
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3122
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3135
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3136
		Category: hMisc,
		//line sql.y: 3137
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3138
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3154
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3155
		Category: hMisc,
		//line sql.y: 3156
		Text: `SHOW JOBS
`,
		//line sql.y: 3157
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3165
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3166
		Category: hMisc,
		//line sql.y: 3167
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3169
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3192
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3193
		Category: hMisc,
		//line sql.y: 3194
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3195
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3211
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3212
		Category: hDDL,
		//line sql.y: 3213
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3214
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3240
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3241
		Category: hDDL,
		//line sql.y: 3242
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3254
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3255
		Category: hMisc,
		//line sql.y: 3256
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3265
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3266
		Category: hCfg,
		//line sql.y: 3267
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3268
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3287
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3288
		Category: hDDL,
		//line sql.y: 3289
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3290
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3308
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3309
		Category: hPriv,
		//line sql.y: 3310
		Text: `SHOW USERS
`,
		//line sql.y: 3311
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3319
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3320
		Category: hPriv,
		//line sql.y: 3321
		Text: `SHOW ROLES
`,
		//line sql.y: 3322
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3367
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3368
		Category: hMisc,
		//line sql.y: 3369
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3605
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3606
		Category: hMisc,
		//line sql.y: 3607
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3610
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3627
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3628
		Category: hDDL,
		//line sql.y: 3629
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3656
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4190
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4191
		Category: hDDL,
		//line sql.y: 4192
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4202
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4256
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4257
		Category: hDML,
		//line sql.y: 4258
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4259
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4267
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4268
		Category: hPriv,
		//line sql.y: 4269
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4270
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4292
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4293
		Category: hPriv,
		//line sql.y: 4294
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4295
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4313
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4314
		Category: hDDL,
		//line sql.y: 4315
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4316
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4349
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4350
		Category: hDDL,
		//line sql.y: 4351
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4359
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4573
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4574
		Category: hTxn,
		//line sql.y: 4575
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4576
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4584
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4585
		Category: hMisc,
		//line sql.y: 4586
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4589
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4606
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4607
		Category: hTxn,
		//line sql.y: 4608
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4609
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4624
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4625
		Category: hTxn,
		//line sql.y: 4626
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4634
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4647
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4648
		Category: hTxn,
		//line sql.y: 4649
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4652
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4676
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4677
		Category: hTxn,
		//line sql.y: 4678
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4679
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4792
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4793
		Category: hDDL,
		//line sql.y: 4794
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4795
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4864
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4865
		Category: hDML,
		//line sql.y: 4866
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4871
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4890
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4891
		Category: hDML,
		//line sql.y: 4892
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4896
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5001
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5002
		Category: hDML,
		//line sql.y: 5003
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5010
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5184
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5185
		Category: hDML,
		//line sql.y: 5186
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5197
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5198
		Category: hDML,
		//line sql.y: 5199
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 5211
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5286
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5287
		Category: hDML,
		//line sql.y: 5288
		Text: `TABLE <tablename>
`,
		//line sql.y: 5289
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5562
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5563
		Category: hDML,
		//line sql.y: 5564
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5565
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5655
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5656
		Category: hDML,
		//line sql.y: 5657
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5675
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
