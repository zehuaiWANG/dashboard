// Copyright 2017 The Kubernetes Authors.
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

// // Copyright 2017 The Kubernetes Authors.
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
//
// import {Component, Input, OnInit} from '@angular/core';
// import {MatTableDataSource} from '@angular/material';
// import {RoleRef} from 'typings/backendapi';
//
// @Component({
//   selector: 'kd-role-ref-list',
//   templateUrl: './template.html',
// })
// export class RoleRefComponent implements OnInit {
//   @Input() initialized: boolean;
//   @Input() roleRef: RoleRef;
//
//   ngOnInit(): void {
//
//   }
//
//   getSubjectColumns(): string[] {
//     return ['apiGroup', 'kind', 'name'];
//   }
//
//   getDataSource(): MatTableDataSource<RoleRef> {
//     const tableData = new MatTableDataSource<RoleRef>();
//     tableData.data = this.roleRef;
//
//     return tableData;
//   }
// }
