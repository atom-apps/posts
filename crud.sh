#/bin/bash

table=article_payment
title=文章
module=posts

# atomctl gen crud $table $module  --backend ../../opensource/backend --force --route /v1/$module/$table --title=$title --module-title=内容管理 # --only-backend 
atomctl gen crud $table $module --route /v1/$module/articles/{article_id}/dig --title=$title --module-title=内容管理 # --only-backend 