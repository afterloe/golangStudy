/*
 Target Server Type    : PostgreSQL
 Target Server Version : 90405
 File Encoding         : 65001

 Date: 28/02/2018 15:39:14
*/


-- ----------------------------
-- Table structure for administrative_area
-- ----------------------------
DROP TABLE IF EXISTS "administrative_area";
CREATE TABLE "administrative_area" (
  "id" int8 NOT NULL DEFAULT NULL,
  "level" int8 DEFAULT NULL,
  "basename" varchar(255) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "code" varchar(255) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "name" varchar(255) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "parent_id" varchar(255) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
ALTER TABLE "administrative_area" OWNER TO "ascs";

-- ----------------------------
-- Primary Key structure for table administrative_area
-- ----------------------------
ALTER TABLE "administrative_area" ADD CONSTRAINT "administrative_area_pkey" PRIMARY KEY ("id");
