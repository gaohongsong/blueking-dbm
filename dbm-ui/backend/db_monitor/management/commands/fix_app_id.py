# -*- coding: utf-8 -*-
"""
TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-DB管理系统(BlueKing-BK-DBM) available.
Copyright (C) 2017-2023 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at https://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""
import logging

from django.core.management.base import BaseCommand

from backend import env
from backend.components import CCApi
from backend.exceptions import ApiResultError

logger = logging.getLogger("root")


class Command(BaseCommand):
    help = "实例标签app_id改为appid"

    def handle(self, *args, **options):
        instances = CCApi.list_service_instance(
            {"bk_biz_id": env.DBA_APP_BK_BIZ_ID, "page": {"start": 0, "limit": 500}}
        )["info"]
        for instance in instances:
            try:
                CCApi.add_label_for_service_instance(
                    {
                        "bk_biz_id": env.DBA_APP_BK_BIZ_ID,
                        "instance_ids": [instance["id"]],
                        "labels": {"appid": instance["labels"]["app_id"]},
                    }
                )
            except ApiResultError as e:
                logger.error("add_label_for_service_instance bk_instance_ids: %s, error: %s", instance, e)
