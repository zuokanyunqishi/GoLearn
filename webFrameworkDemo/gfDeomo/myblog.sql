/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 127.0.0.1:3306
 Source Schema         : myblog

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 12/07/2022 00:03:24
*/


create database blog;
use blog;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_tags
-- ----------------------------
DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags`
(
    `article_id` int unsigned NOT NULL,
    `tag_id`     int unsigned NOT NULL,
    KEY `article_tags_article_id_index` (`article_id`),
    KEY `article_tags_tag_id_index` (`tag_id`),
    CONSTRAINT `article_tags_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    CONSTRAINT `article_tags_tag_id_foreign` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of article_tags
-- ----------------------------
BEGIN;
INSERT INTO `article_tags` (`article_id`, `tag_id`)
VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`
(
    `id`            int unsigned                                               NOT NULL AUTO_INCREMENT,
    `title`         varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL COMMENT '文章标题',
    `keyword`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'keywords',
    `desc`          varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL COMMENT '描述',
    `content`       longtext CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci COMMENT '文章内容,markdown格式',
    `user_id`       int                                                        NOT NULL DEFAULT '0' COMMENT '文章编写人,对应users表',
    `cate_id`       int                                                        NOT NULL DEFAULT '0' COMMENT '文章分类',
    `comment_count` int                                                        NOT NULL DEFAULT '0' COMMENT '评论数量',
    `read_count`    int                                                        NOT NULL DEFAULT '0' COMMENT '阅读数量',
    `status`        tinyint                                                    NOT NULL DEFAULT '0' COMMENT '文章状态:0-公开;1-私密',
    `sort`          int                                                        NOT NULL DEFAULT '0' COMMENT '排序',
    `created_at`    timestamp                                                  NULL     DEFAULT NULL,
    `updated_at`    timestamp                                                  NULL     DEFAULT NULL,
    `html_content`  longtext CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci COMMENT '文章内容,html格式',
    `list_pic`      longtext CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci COMMENT '文章列表图',
    PRIMARY KEY (`id`),
    KEY `articles_cate_id_index` (`cate_id`),
    KEY `articles_user_id_index` (`user_id`),
    KEY `articles_title_index` (`title`),
    KEY `articles_created_at_index` (`created_at`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of articles
-- ----------------------------
BEGIN;
INSERT INTO `articles` (`id`, `title`, `keyword`, `desc`, `content`, `user_id`, `cate_id`, `comment_count`,
                        `read_count`, `status`, `sort`, `created_at`, `updated_at`, `html_content`, `list_pic`)
VALUES (1, 'mysql基础', '', '学习mysql',
        '# mysql入门\r\n* 查看表结构\r\n\r\n> ```sql\r\n>   desc table;\r\n>  ```\r\n\r\n| Field   | Type             | Null   | Key   | Default   | Extra          |\r\n|---------|------------------|--------|-------|-----------|----------------|\r\n| id      | int(10) unsigned | NO     | PRI   | <null>    | auto_increment |\r\n| address | varchar(30)      | NO     |       |           |                ||\r\n\r\n\r\n*  查看表创建语句\r\n	```sql\r\n	 show create TABLE yunnan_answer_chancelog\r\n	  yunnan_answer_chancelog | CREATE TABLE `yunnan_answer_chancelog` (\r\n	  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,\r\n	  `uuid` char(16) NOT NULL DEFAULT \'\' COMMENT \'uuid\',\r\n	  `mobile` char(11) NOT NULL DEFAULT \'\' COMMENT \'手机号\',\r\n	  `state` tinyint(1) unsigned NOT NULL COMMENT \'加机会状态，1：成功，2：失败\',\r\n	  `ctime` datetime NOT NULL DEFAULT \'0001-01-01 00:00:00\' COMMENT \'增加机会时间\',\r\n	  `type` tinyint(4) NOT NULL DEFAULT \'0\' COMMENT \'增加机会的类型,1是每天赠送机会,2是答题增加的机会\',\r\n	  PRIMARY KEY (`id`),\r\n	  KEY `mobile` (`mobile`,`ctime`)\r\n	) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8 COMMENT=\'(云南4月答题摇奖)添加机会记录\';\r\n	```\r\n* 修改表名\r\n	 ```sql\r\n	  alter table old_name rename [to] new_name; \r\n	 ```\r\n* 修改表字段数据类型\r\n	 ```sql\r\n	  #alter table <table_name> modify <字段名> <数据类型>\r\n	  alter table anshan_article modify user varchar(20)\r\n	 ```\r\n* 修改表字段名\r\n	```sql\r\n	  alter table <表名> change <旧字段名> <新字段名> <新数据类型>\r\n	  alter table anshan_article change user_id use_info char(15)\r\n	```\r\n* 添加字段\r\n	```sql\r\n	   alter table <表名> add <新字段名> <数据类型>\r\n	   alter table tb  add column1 varchar(12) not null;\r\n	   #某个字段后添加\r\n		alter table tb  add column1 varchar(12) not null first[after name];\r\n	```\r\n* 删除字段\r\n	```sql\r\n	   alter table <table_name> drop <field_name>\r\n	```\r\n* 修改表的存储引擎\r\n	```sql\r\n	   alter table <table_name> engine= <change_name>\r\n	```\r\n* 删除表\r\n	```sql\r\n	   drop table [if exists] <table_name1>,<table_name2>\r\n	```\r\n\r\n##### 数据类型\r\n*  整数数据类型\r\n\r\n|type |min|max|length|desc|\r\n|-----|-----|-----|------|------|\r\n|tinyint|    -128~127|0~255|1个字节|很小的整数|\r\n|smallint|32768~32767|0~65535|2个字节|小的整数|\r\n|mediumint|-8388608~8388607|0~16777215|3个字节|中等的大小整数|\r\n|int|-2147483648~2147483647|0~4294967295|4个字节|普通大小整数|\r\n|bigint|-9223372036854775808~9223372036854775807|0~18446744073709551615|8个字节|大整数|\r\n* 浮点数据类型\r\n\r\n|type |min|max|length|desc|\r\n|-----|-----|-----|------|------|\r\n|float|-3.402823466E+38~-1.175494351E-38|0和 1.175494351E-38~3.402823466E+38|4个字节|单精度浮点数|\r\n|double|-1.7976931348623157E+308~-2.2250738585072014E-308|0和2.2250738585072014E-308~1.7976931348623157E+308|8个字节|双精度浮点数|\r\n|decimal(m,d),dec|||M+2个字节|压缩的\'严格\'定点数|\r\n\r\n* 日期和时间数据类型\r\n\r\n|type |日期格式|日期范围|length|desc|\r\n|-----|-----|-----|------|------|\r\n|yaer|YYYY|1901~2155|1字节|年|\r\n|time|HH:MM:SS|-838:59:59 ~ 838:59:59|3字节|时分秒|\r\n|date|YYYY-MM-DD|1000-01-01 ~ 9999-12-3|3字节|年月日|\r\n|datetime|YYYY-MM-DD HH:MM:SS|1000-01-01 00:00:00 ~ 9999-12-31 23:59:59|8字节|年月日时分秒|\r\n|timestamp|YYYY-MM-DD HH:MM:SS|1970-01-01 00:00:00 UTC ~ 2038-01-19 03:14:07 UTC|4四节|时间戳|\r\n* 文本字符串类型\r\n\r\n|type |length(字节)|desc|\r\n|-----|-----|------|\r\n|char(M)|1<-M<=255|固定长度非二进制字符串|\r\n|varchar(M)|L+1,L<=M 和 1<=M<=255||\r\n|tinytext||||\r\n',
        1, 1, 1, 8, 0, 0, '2018-06-02 11:23:47', '2018-06-09 19:42:57',
        '<h1 id=\"h1-mysql-\"><a name=\"mysql入门\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>mysql入门</h1><ul>\r\n<li>查看表结构</li></ul>\r\n<blockquote>\r\n<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">  desc table</span><span class=\"pun\">;</span></code></li></ol></pre>\r\n</blockquote>\r\n<table>\r\n<thead>\r\n<tr>\r\n<th>Field</th>\r\n<th>Type</th>\r\n<th>Null</th>\r\n<th>Key</th>\r\n<th>Default</th>\r\n<th>Extra</th>\r\n</tr>\r\n</thead>\r\n<tbody>\r\n<tr>\r\n<td>id</td>\r\n<td>int(10) unsigned</td>\r\n<td>NO</td>\r\n<td>PRI</td>\r\n<td>&lt;null&gt;</td>\r\n<td>auto_increment</td>\r\n</tr>\r\n<tr>\r\n<td>address</td>\r\n<td>varchar(30)</td>\r\n<td>NO</td>\r\n<td></td>\r\n<td></td>\r\n<td></td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<ul>\r\n<li>查看表创建语句<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">  show create TABLE yunnan_answer_chancelog</span></code></li><li class=\"L1\"><code class=\"lang-sql\"><span class=\"pln\">   yunnan_answer_chancelog </span><span class=\"pun\">|</span><span class=\"pln\"> CREATE TABLE </span><span class=\"str\">`yunnan_answer_chancelog`</span><span class=\"pln\"> </span><span class=\"pun\">(</span></code></li><li class=\"L2\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`id`</span><span class=\"pln\"> </span><span class=\"kwd\">int</span><span class=\"pun\">(</span><span class=\"lit\">11</span><span class=\"pun\">)</span><span class=\"pln\"> </span><span class=\"kwd\">unsigned</span><span class=\"pln\"> NOT NULL AUTO_INCREMENT</span><span class=\"pun\">,</span></code></li><li class=\"L3\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`uuid`</span><span class=\"pln\"> </span><span class=\"kwd\">char</span><span class=\"pun\">(</span><span class=\"lit\">16</span><span class=\"pun\">)</span><span class=\"pln\"> NOT NULL DEFAULT </span><span class=\"str\">\'\'</span><span class=\"pln\"> COMMENT </span><span class=\"str\">\'uuid\'</span><span class=\"pun\">,</span></code></li><li class=\"L4\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`mobile`</span><span class=\"pln\"> </span><span class=\"kwd\">char</span><span class=\"pun\">(</span><span class=\"lit\">11</span><span class=\"pun\">)</span><span class=\"pln\"> NOT NULL DEFAULT </span><span class=\"str\">\'\'</span><span class=\"pln\"> COMMENT </span><span class=\"str\">\'手机号\'</span><span class=\"pun\">,</span></code></li><li class=\"L5\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`state`</span><span class=\"pln\"> tinyint</span><span class=\"pun\">(</span><span class=\"lit\">1</span><span class=\"pun\">)</span><span class=\"pln\"> </span><span class=\"kwd\">unsigned</span><span class=\"pln\"> NOT NULL COMMENT </span><span class=\"str\">\'加机会状态，1：成功，2：失败\'</span><span class=\"pun\">,</span></code></li><li class=\"L6\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`ctime`</span><span class=\"pln\"> datetime NOT NULL DEFAULT </span><span class=\"str\">\'0001-01-01 00:00:00\'</span><span class=\"pln\"> COMMENT </span><span class=\"str\">\'增加机会时间\'</span><span class=\"pun\">,</span></code></li><li class=\"L7\"><code class=\"lang-sql\"><span class=\"pln\">   </span><span class=\"str\">`type`</span><span class=\"pln\"> tinyint</span><span class=\"pun\">(</span><span class=\"lit\">4</span><span class=\"pun\">)</span><span class=\"pln\"> NOT NULL DEFAULT </span><span class=\"str\">\'0\'</span><span class=\"pln\"> COMMENT </span><span class=\"str\">\'增加机会的类型,1是每天赠送机会,2是答题增加的机会\'</span><span class=\"pun\">,</span></code></li><li class=\"L8\"><code class=\"lang-sql\"><span class=\"pln\">   PRIMARY KEY </span><span class=\"pun\">(</span><span class=\"str\">`id`</span><span class=\"pun\">),</span></code></li><li class=\"L9\"><code class=\"lang-sql\"><span class=\"pln\">   KEY </span><span class=\"str\">`mobile`</span><span class=\"pln\"> </span><span class=\"pun\">(</span><span class=\"str\">`mobile`</span><span class=\"pun\">,</span><span class=\"str\">`ctime`</span><span class=\"pun\">)</span></code></li><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\"> </span><span class=\"pun\">)</span><span class=\"pln\"> ENGINE</span><span class=\"pun\">=</span><span class=\"typ\">InnoDB</span><span class=\"pln\"> AUTO_INCREMENT</span><span class=\"pun\">=</span><span class=\"lit\">33</span><span class=\"pln\"> DEFAULT CHARSET</span><span class=\"pun\">=</span><span class=\"pln\">utf8 COMMENT</span><span class=\"pun\">=</span><span class=\"str\">\'(云南4月答题摇奖)添加机会记录\'</span><span class=\"pun\">;</span></code></li></ol></pre>\r\n</li><li>修改表名<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">    alter table old_name rename </span><span class=\"pun\">[</span><span class=\"pln\">to</span><span class=\"pun\">]</span><span class=\"pln\"> new_name</span><span class=\"pun\">;</span></code></li></ol></pre>\r\n</li><li>修改表字段数据类型<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">    </span><span class=\"com\">#alter table &lt;table_name&gt; modify &lt;字段名&gt; &lt;数据类型&gt;</span></code></li><li class=\"L1\"><code class=\"lang-sql\"><span class=\"pln\">    alter table anshan_article modify user varchar</span><span class=\"pun\">(</span><span class=\"lit\">20</span><span class=\"pun\">)</span></code></li></ol></pre>\r\n</li><li>修改表字段名<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">    alter table </span><span class=\"pun\">&lt;表名&gt;</span><span class=\"pln\"> change </span><span class=\"pun\">&lt;旧字段名&gt;</span><span class=\"pln\"> </span><span class=\"pun\">&lt;新字段名&gt;</span><span class=\"pln\"> </span><span class=\"pun\">&lt;新数据类型&gt;</span></code></li><li class=\"L1\"><code class=\"lang-sql\"><span class=\"pln\">    alter table anshan_article change user_id use_info </span><span class=\"kwd\">char</span><span class=\"pun\">(</span><span class=\"lit\">15</span><span class=\"pun\">)</span></code></li></ol></pre>\r\n</li><li>添加字段<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">     alter table </span><span class=\"pun\">&lt;表名&gt;</span><span class=\"pln\"> add </span><span class=\"pun\">&lt;新字段名&gt;</span><span class=\"pln\"> </span><span class=\"pun\">&lt;数据类型&gt;</span></code></li><li class=\"L1\"><code class=\"lang-sql\"><span class=\"pln\">     alter table tb  add column1 varchar</span><span class=\"pun\">(</span><span class=\"lit\">12</span><span class=\"pun\">)</span><span class=\"pln\"> </span><span class=\"kwd\">not</span><span class=\"pln\"> </span><span class=\"kwd\">null</span><span class=\"pun\">;</span></code></li><li class=\"L2\"><code class=\"lang-sql\"><span class=\"pln\">     </span><span class=\"com\">#某个字段后添加</span></code></li><li class=\"L3\"><code class=\"lang-sql\"><span class=\"pln\">      alter table tb  add column1 varchar</span><span class=\"pun\">(</span><span class=\"lit\">12</span><span class=\"pun\">)</span><span class=\"pln\"> </span><span class=\"kwd\">not</span><span class=\"pln\"> </span><span class=\"kwd\">null</span><span class=\"pln\"> first</span><span class=\"pun\">[</span><span class=\"pln\">after name</span><span class=\"pun\">];</span></code></li></ol></pre>\r\n</li><li>删除字段<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">     alter table </span><span class=\"str\">&lt;table_name&gt;</span><span class=\"pln\"> drop </span><span class=\"str\">&lt;field_name&gt;</span></code></li></ol></pre>\r\n</li><li>修改表的存储引擎<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">     alter table </span><span class=\"str\">&lt;table_name&gt;</span><span class=\"pln\"> engine</span><span class=\"pun\">=</span><span class=\"pln\"> </span><span class=\"str\">&lt;change_name&gt;</span></code></li></ol></pre>\r\n</li><li>删除表<pre class=\"prettyprint linenums prettyprinted\" style=\"\"><ol class=\"linenums\"><li class=\"L0\"><code class=\"lang-sql\"><span class=\"pln\">     drop table </span><span class=\"pun\">[</span><span class=\"kwd\">if</span><span class=\"pln\"> exists</span><span class=\"pun\">]</span><span class=\"pln\"> </span><span class=\"str\">&lt;table_name1&gt;</span><span class=\"pun\">,&lt;</span><span class=\"pln\">table_name2</span><span class=\"pun\">&gt;</span></code></li></ol></pre>\r\n</li></ul>\r\n<h5 id=\"h5-u6570u636Eu7C7Bu578B\"><a name=\"数据类型\" class=\"reference-link\"></a><span class=\"header-link octicon octicon-link\"></span>数据类型</h5><ul>\r\n<li>整数数据类型</li></ul>\r\n<table>\r\n<thead>\r\n<tr>\r\n<th>type</th>\r\n<th>min</th>\r\n<th>max</th>\r\n<th>length</th>\r\n<th>desc</th>\r\n</tr>\r\n</thead>\r\n<tbody>\r\n<tr>\r\n<td>tinyint</td>\r\n<td>-128~127</td>\r\n<td>0~255</td>\r\n<td>1个字节</td>\r\n<td>很小的整数</td>\r\n</tr>\r\n<tr>\r\n<td>smallint</td>\r\n<td>32768~32767</td>\r\n<td>0~65535</td>\r\n<td>2个字节</td>\r\n<td>小的整数</td>\r\n</tr>\r\n<tr>\r\n<td>mediumint</td>\r\n<td>-8388608~8388607</td>\r\n<td>0~16777215</td>\r\n<td>3个字节</td>\r\n<td>中等的大小整数</td>\r\n</tr>\r\n<tr>\r\n<td>int</td>\r\n<td>-2147483648~2147483647</td>\r\n<td>0~4294967295</td>\r\n<td>4个字节</td>\r\n<td>普通大小整数</td>\r\n</tr>\r\n<tr>\r\n<td>bigint</td>\r\n<td>-9223372036854775808~9223372036854775807</td>\r\n<td>0~18446744073709551615</td>\r\n<td>8个字节</td>\r\n<td>大整数</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<ul>\r\n<li>浮点数据类型</li></ul>\r\n<table>\r\n<thead>\r\n<tr>\r\n<th>type</th>\r\n<th>min</th>\r\n<th>max</th>\r\n<th>length</th>\r\n<th>desc</th>\r\n</tr>\r\n</thead>\r\n<tbody>\r\n<tr>\r\n<td>float</td>\r\n<td>-3.402823466E+38~-1.175494351E-38</td>\r\n<td>0和 1.175494351E-38~3.402823466E+38</td>\r\n<td>4个字节</td>\r\n<td>单精度浮点数</td>\r\n</tr>\r\n<tr>\r\n<td>double</td>\r\n<td>-1.7976931348623157E+308~-2.2250738585072014E-308</td>\r\n<td>0和2.2250738585072014E-308~1.7976931348623157E+308</td>\r\n<td>8个字节</td>\r\n<td>双精度浮点数</td>\r\n</tr>\r\n<tr>\r\n<td>decimal(m,d),dec</td>\r\n<td></td>\r\n<td></td>\r\n<td>M+2个字节</td>\r\n<td>压缩的’严格’定点数</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<ul>\r\n<li>日期和时间数据类型</li></ul>\r\n<table>\r\n<thead>\r\n<tr>\r\n<th>type</th>\r\n<th>日期格式</th>\r\n<th>日期范围</th>\r\n<th>length</th>\r\n<th>desc</th>\r\n</tr>\r\n</thead>\r\n<tbody>\r\n<tr>\r\n<td>yaer</td>\r\n<td>YYYY</td>\r\n<td>1901~2155</td>\r\n<td>1字节</td>\r\n<td>年</td>\r\n</tr>\r\n<tr>\r\n<td>time</td>\r\n<td>HH:MM:SS</td>\r\n<td>-838:59:59 ~ 838:59:59</td>\r\n<td>3字节</td>\r\n<td>时分秒</td>\r\n</tr>\r\n<tr>\r\n<td>date</td>\r\n<td>YYYY-MM-DD</td>\r\n<td>1000-01-01 ~ 9999-12-3</td>\r\n<td>3字节</td>\r\n<td>年月日</td>\r\n</tr>\r\n<tr>\r\n<td>datetime</td>\r\n<td>YYYY-MM-DD HH:MM:SS</td>\r\n<td>1000-01-01 00:00:00 ~ 9999-12-31 23:59:59</td>\r\n<td>8字节</td>\r\n<td>年月日时分秒</td>\r\n</tr>\r\n<tr>\r\n<td>timestamp</td>\r\n<td>YYYY-MM-DD HH:MM:SS</td>\r\n<td>1970-01-01 00:00:00 UTC ~ 2038-01-19 03:14:07 UTC</td>\r\n<td>4四节</td>\r\n<td>时间戳</td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n<ul>\r\n<li>文本字符串类型</li></ul>\r\n<table>\r\n<thead>\r\n<tr>\r\n<th>type</th>\r\n<th>length(字节)</th>\r\n<th>desc</th>\r\n</tr>\r\n</thead>\r\n<tbody>\r\n<tr>\r\n<td>char(M)</td>\r\n<td>1&lt;-M&lt;=255</td>\r\n<td>固定长度非二进制字符串</td>\r\n</tr>\r\n<tr>\r\n<td>varchar(M)</td>\r\n<td>L+1,L&lt;=M 和 1&lt;=M&lt;=255</td>\r\n<td></td>\r\n</tr>\r\n<tr>\r\n<td>tinytext</td>\r\n<td></td>\r\n<td></td>\r\n</tr>\r\n</tbody>\r\n</table>\r\n',
        'http://static.open-open.com/lib/uploadImg/20150120/20150120223218_128.jpg');
COMMIT;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`
(
    `id`         int unsigned                                               NOT NULL AUTO_INCREMENT,
    `parent_id`  int                                                             DEFAULT NULL,
    `lft`        int                                                             DEFAULT NULL,
    `rgt`        int                                                             DEFAULT NULL,
    `depth`      int                                                             DEFAULT NULL,
    `name`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `created_at` timestamp                                                  NULL DEFAULT NULL,
    `updated_at` timestamp                                                  NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `categories_parent_id_index` (`parent_id`),
    KEY `categories_lft_index` (`lft`),
    KEY `categories_rgt_index` (`rgt`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 6
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of categories
-- ----------------------------
BEGIN;
INSERT INTO `categories` (`id`, `parent_id`, `lft`, `rgt`, `depth`, `name`, `created_at`, `updated_at`)
VALUES (1, NULL, 1, 4, 0, '数据库', '2018-06-02 10:49:46', '2018-06-02 11:18:11');
INSERT INTO `categories` (`id`, `parent_id`, `lft`, `rgt`, `depth`, `name`, `created_at`, `updated_at`)
VALUES (2, NULL, 5, 8, 0, '后端', '2018-06-02 10:50:09', '2018-06-02 11:18:11');
INSERT INTO `categories` (`id`, `parent_id`, `lft`, `rgt`, `depth`, `name`, `created_at`, `updated_at`)
VALUES (3, 2, 6, 7, 1, 'php', '2018-06-02 11:14:29', '2018-06-02 11:18:11');
INSERT INTO `categories` (`id`, `parent_id`, `lft`, `rgt`, `depth`, `name`, `created_at`, `updated_at`)
VALUES (4, NULL, 9, 10, 0, 'linux', '2018-06-02 11:17:26', '2018-06-02 11:18:11');
INSERT INTO `categories` (`id`, `parent_id`, `lft`, `rgt`, `depth`, `name`, `created_at`, `updated_at`)
VALUES (5, 1, 2, 3, 1, 'mysql', '2018-06-02 11:18:11', '2018-06-02 11:18:11');
COMMIT;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`
(
    `id`          int unsigned                                       NOT NULL AUTO_INCREMENT,
    `article_id`  int                                                NOT NULL,
    `parent_id`   int                                                NOT NULL DEFAULT '0',
    `content`     text CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `name`        varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `email`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `website`     varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `avatar`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `ip`          varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `city`        varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `user_id`     int                                                NOT NULL DEFAULT '0',
    `target_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci  DEFAULT NULL,
    `created_at`  timestamp                                          NULL     DEFAULT NULL,
    `updated_at`  timestamp                                          NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of comments
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for links
-- ----------------------------
DROP TABLE IF EXISTS `links`;
CREATE TABLE `links`
(
    `id`         int unsigned                                               NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `url`        varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `sequence`   smallint                                                   NOT NULL,
    `created_at` timestamp                                                  NULL DEFAULT NULL,
    `updated_at` timestamp                                                  NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of links
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations`
(
    `migration` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `batch`     int                                                        NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of migrations
-- ----------------------------
BEGIN;
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2014_10_12_000000_create_users_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2014_10_12_100000_create_password_resets_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2017_12_23_085521_add_users_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2017_12_28_111756_create_article_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2017_12_30_142122_create_categories_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2017_12_31_143011_create_tags_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_03_121356_add_html_content_to_articles_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_04_021159_create_article_tags_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_05_062957_create_navigations_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_09_144849_create_systems_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_10_031610_create_link_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_01_10_035756_create_pages_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_03_22_152609_create_comments_table', 1);
INSERT INTO `migrations` (`migration`, `batch`)
VALUES ('2018_05_25_093803_add_list_pic_to_articles_table', 1);
COMMIT;

-- ----------------------------
-- Table structure for navigations
-- ----------------------------
DROP TABLE IF EXISTS `navigations`;
CREATE TABLE `navigations`
(
    `id`              int unsigned                                               NOT NULL AUTO_INCREMENT,
    `name`            varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `url`             varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `state`           tinyint                                                    NOT NULL DEFAULT '0' COMMENT '0-正常显示;1-隐藏',
    `sequence`        smallint                                                   NOT NULL COMMENT '排序',
    `nav_type`        tinyint                                                    NOT NULL DEFAULT '0' COMMENT '导航类型:0-自定义;1-分类导航',
    `article_cate_id` int                                                        NOT NULL DEFAULT '0' COMMENT '文章分类id',
    `created_at`      timestamp                                                  NULL     DEFAULT NULL,
    `updated_at`      timestamp                                                  NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 6
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of navigations
-- ----------------------------
BEGIN;
INSERT INTO `navigations` (`id`, `name`, `url`, `state`, `sequence`, `nav_type`, `article_cate_id`, `created_at`,
                           `updated_at`)
VALUES (1, 'php', 'http://mynewblog.cn/category/3', 0, 0, 1, 3, '2018-06-02 11:47:34', '2018-06-02 11:47:34');
INSERT INTO `navigations` (`id`, `name`, `url`, `state`, `sequence`, `nav_type`, `article_cate_id`, `created_at`,
                           `updated_at`)
VALUES (2, '后端', 'http://mynewblog.cn/category/2', 0, 0, 1, 2, '2018-06-02 11:47:47', '2018-06-02 11:47:47');
INSERT INTO `navigations` (`id`, `name`, `url`, `state`, `sequence`, `nav_type`, `article_cate_id`, `created_at`,
                           `updated_at`)
VALUES (3, '数据库', 'http://mynewblog.cn/category/1', 0, 0, 1, 1, '2018-06-02 11:47:50', '2018-06-02 11:47:50');
INSERT INTO `navigations` (`id`, `name`, `url`, `state`, `sequence`, `nav_type`, `article_cate_id`, `created_at`,
                           `updated_at`)
VALUES (5, '时间线', 'http://mynewblog.cn/article/select', 0, 0, 0, 0, '2018-06-02 12:04:22', '2018-06-02 12:04:22');
COMMIT;

-- ----------------------------
-- Table structure for pages
-- ----------------------------
DROP TABLE IF EXISTS `pages`;
CREATE TABLE `pages`
(
    `id`           int unsigned                                               NOT NULL AUTO_INCREMENT,
    `title`        varchar(200) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '页面标题',
    `link_alias`   varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci          DEFAULT NULL COMMENT '链接别名',
    `keyword`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '关键词',
    `desc`         varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
    `content`      longtext CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci COMMENT '页面markdown格式',
    `html_content` longtext CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci COMMENT '页面html',
    `created_at`   timestamp                                                  NULL     DEFAULT NULL,
    `updated_at`   timestamp                                                  NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `pages_link_alias_unique` (`link_alias`),
    KEY `pages_link_alias_index` (`link_alias`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of pages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for password_resets
-- ----------------------------
DROP TABLE IF EXISTS `password_resets`;
CREATE TABLE `password_resets`
(
    `email`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `token`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `created_at` timestamp                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `password_resets_email_index` (`email`),
    KEY `password_resets_token_index` (`token`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of password_resets
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for systems
-- ----------------------------
DROP TABLE IF EXISTS `systems`;
CREATE TABLE `systems`
(
    `id`    int unsigned                                               NOT NULL AUTO_INCREMENT,
    `key`   varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `value` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `systems_key_unique` (`key`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 12
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of systems
-- ----------------------------
BEGIN;
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (1, 'blog_name', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (2, 'motto', '天空任鸟飞');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (3, 'title', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (4, 'seo_keyword', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (5, 'seo_desc', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (6, 'github_url', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (7, 'qq', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (8, 'comment_plugin', 'changyan');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (9, 'icp', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (10, 'statistics_script', '');
INSERT INTO `systems` (`id`, `key`, `value`)
VALUES (11, 'comment_script', '');
COMMIT;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`
(
    `id`         int unsigned                                               NOT NULL AUTO_INCREMENT,
    `tag_name`   varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL COMMENT '标签名字',
    `created_at` timestamp                                                  NULL DEFAULT NULL,
    `updated_at` timestamp                                                  NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `tags_tag_name_unique` (`tag_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of tags
-- ----------------------------
BEGIN;
INSERT INTO `tags` (`id`, `tag_name`, `created_at`, `updated_at`)
VALUES (1, 'mysql', '2018-06-02 11:23:47', '2018-06-02 11:23:47');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`             int unsigned                                               NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `email`          varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `password`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL,
    `remember_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci          DEFAULT NULL,
    `created_at`     timestamp                                                  NULL     DEFAULT NULL,
    `updated_at`     timestamp                                                  NULL     DEFAULT NULL,
    `user_pic`       varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '博客用户头像',
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_email_unique` (`email`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb3
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
