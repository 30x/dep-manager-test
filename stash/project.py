# Copyright 2017 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Creates a single project with specified service accounts and APIs enabled."""


def GenerateConfig(context):
  """Generates config."""

  project_id = context.env['name']
  region = context.properties['region']
  billing_name = 'billing_' + project_id

  resources = [{
      'name': project_id,
      'type': 'cloudresourcemanager.v1.project',
      'properties': {
          'projectId': project_id,
          'labels': context.properties['labels'],
          'parent': {
              'type': 'organization',
              'id': context.properties['organization-id']
          }
      },
      'accessControl': {
          'gcpIamPolicy': context.properties['iam-policy']
      }
  }, {
      'name': billing_name,
      'type': 'deploymentmanager.v2.virtual.projectBillingInfo',
      'metadata': {
          'dependsOn': [project_id]
      },
      'properties': {
          'name': 'projects/' + project_id,
          'billingAccountName': context.properties['billing-account-name']
      }
  }, {
      'name': 'apis',
      'type': 'apis.py',
      'properties': {
          'project': project_id,
          'billing': billing_name,
          'apis': context.properties['apis']
      }
  }, {
      'name': 'service-accounts',
      'type': 'service-accounts.py',
      'properties': {
          'project': project_id,
          'service-accounts': context.properties['service-accounts']
      }
  },
#   {
#       'name': 'appengine-app',
#       'type': 'appengine.v1.app',
#       'properties': {
#           'id': project_id,
#           'location_id': context.properties['region']
#       }
#   },
#   {
#       'name': 'appengine-version',
#       'type': 'appengine.v1.version',
#       'properties': {  
#         'appsId': project_id,
#         'servicesId': 'default',
#         'deployment': {
#             'files': {
#                 'my-resource-file1': {
#                     'sourceUrl': "https://storage.googleapis.com/py-test/main.py"
#                 },
#             }, 
#         },
#         'id': '1',
#         'runtime': 'python27',
#         'handlers': [{
#             'urlRegex': '/*',
#             'script': {
#                 'scriptPath': 'default.app'
#             }
#         }]
#      }
#   }
  ]

  return {'resources': resources}