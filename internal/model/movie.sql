CREATE TABLE `douban_movie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT '' COMMENT '标题',
  `subtitle` varchar(255) DEFAULT '' COMMENT '副标题',
  `other` varchar(255) DEFAULT '' COMMENT '其他',
  `url` varchar(255) DEFAULT '' COMMENT '链接',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `year` varchar(255) DEFAULT '0' COMMENT '年份',
  `area` varchar(255) DEFAULT '' COMMENT '地区',
  `tag` varchar(255) DEFAULT '' COMMENT '标签',
  `star` varchar(10) DEFAULT '0' COMMENT 'star',
  `comment` int(10) unsigned DEFAULT '0' COMMENT '评分',
  `view_number` varchar(255) DEFAULT '0' COMMENT '浏览量',
  `quote` varchar(255) DEFAULT '' COMMENT '引用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='movieTop250';