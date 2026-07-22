
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { ForeignExchangeRatesSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await ForeignExchangeRatesSDK.test()
    equal(null !== testsdk, true)
  })

})
