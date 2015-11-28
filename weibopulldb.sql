CREATE TABLE "crawbuffer" ("uid" TEXT,"time" INTEGER,"crawinstanceuuid" TEXT,"firedurl" TEXT,"ctx" TEXT,"pageno" INTEGER,"Nextpagefound" BOOL);
CREATE TABLE "crawinstance" ("crawinstanceuuid" TEXT, "initializedtime" INTEGER, "crawas" TEXT, "crawcookie" TEXT, "crawnote" TEXT, "crawua" TEXT);
CREATE TABLE "crawresult" ("uid" TEXT,"lasttime" INTEGER DEFAULT (null) ,"lastcrawedpage" INTEGER DEFAULT (0) );
CREATE TABLE "programma_configure" ("confname" TEXT, "cfncval" );
CREATE TABLE "weibocrawtarget" ("typeofowner" TEXT,"uid" TEXT);
CREATE TABLE "weibofeeds" ("uid" TEXT, "textw" TEXT, "page" INTEGER, "natpage" INTEGER, "crawinstanceuuid" TEXT);
