
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { ForeignExchangeRatesSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('ConvertEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when FOREIGNEXCHANGERATES_TEST_LIVE=TRUE.
  afterEach(liveDelay('FOREIGNEXCHANGERATES_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = ForeignExchangeRatesSDK.test()
    const ent = testsdk.Convert()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.FOREIGN_EXCHANGE_RATES_TEST_LIVE
    for (const op of ['create', 'list']) {
      if (maybeSkipControl(t, 'entityOp', 'convert.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set FOREIGN_EXCHANGE_RATES_TEST_CONVERT_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const convert_ref01_ent = client.Convert()
    let convert_ref01_data = setup.data.new.convert['convert_ref01']
    convert_ref01_data['amount'] = setup.idmap['amount01']
    convert_ref01_data['from'] = setup.idmap['from01']
    convert_ref01_data['to'] = setup.idmap['to01']

    convert_ref01_data = await convert_ref01_ent.create(convert_ref01_data)
    assert(null != convert_ref01_data)


    // LIST
    const convert_ref01_match: any = {}
    convert_ref01_match['amount'] = setup.idmap['amount01']
    convert_ref01_match['from'] = setup.idmap['from01']
    convert_ref01_match['to'] = setup.idmap['to01']

    const convert_ref01_list = await convert_ref01_ent.list(convert_ref01_match)

    assert(!isempty(select(convert_ref01_list, { id: convert_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/convert/ConvertTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = ForeignExchangeRatesSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['convert01','convert02','convert03','convert01','convert02','convert03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['FOREIGN_EXCHANGE_RATES_TEST_CONVERT_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'FOREIGN_EXCHANGE_RATES_TEST_CONVERT_ENTID': idmap,
    'FOREIGN_EXCHANGE_RATES_TEST_LIVE': 'FALSE',
    'FOREIGN_EXCHANGE_RATES_TEST_EXPLAIN': 'FALSE',
    'FOREIGN_EXCHANGE_RATES_APIKEY': 'NONE',
  })

  idmap = env['FOREIGN_EXCHANGE_RATES_TEST_CONVERT_ENTID']

  const live = 'TRUE' === env.FOREIGN_EXCHANGE_RATES_TEST_LIVE

  if (live) {
    client = new ForeignExchangeRatesSDK(merge([
      {
        apikey: env.FOREIGN_EXCHANGE_RATES_APIKEY,
      },
      extra
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.FOREIGN_EXCHANGE_RATES_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
