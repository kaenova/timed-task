import React, { useState } from 'react'
import CheckoutItem from './checkout_item'
import axios from 'axios'
import { useQuery } from "@tanstack/react-query"

function Checkout({ className }) {

  const { status, data, error, isFetching } = useQuery(
    ['checkouts'],
    async () => {
      const res = await axios.get("http://127.0.0.1:3001/checkout")
      return res.data
    },
    {
      refetchInterval: 3000,
    },
  )

  if (status == "loading") {
    return (
      <div className={className}>
        <div className="card bg-base-200 shadow-xl">
          <div className="card-body">
            <h1 className='font-bold text-center text-xl mb-3'>Loading</h1>
          </div>
        </div>
      </div>
    )
  }

  if (status == "error") {
    return (
      <div className={className}>
        <div className="card bg-base-200 shadow-xl">
          <div className="card-body">
            <h1 className='font-bold text-center text-xl mb-3'>Error when fetching data</h1>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className={className}>
      <div className="card bg-base-200 shadow-xl">
        <div className="card-body">
          <h1 className='font-bold text-center text-xl mb-3'>Current Checkout</h1>
          {
            data.map((checkout, i) => {
              return < CheckoutItem key={i} data={checkout} /> 
            })
          }
        </div>
      </div>
    </div>
  )
}

export default Checkout