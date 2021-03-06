/*
 * Copyright 2020 Verizon Media
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Add some helpful assertions for `react-testing-library`
import '@testing-library/jest-dom/extend-expect';
import jsdom from 'jsdom';
import { createSerializer } from '@emotion/jest';

const dom = new jsdom.JSDOM('<!doctype html><html><body></body></html>');

expect.addSnapshotSerializer(createSerializer());

// new lines
global.Node = dom.window.Node;
require('mutationobserver-shim');
global.MutationObserver = global.window.MutationObserver;

process.env.APP_ENV = 'unittest';
