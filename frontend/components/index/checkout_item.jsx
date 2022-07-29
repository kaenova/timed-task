import React from 'react'
import { BsXCircle } from 'react-icons/bs'
import TimeBadge from './time_badge'
import axios from 'axios'

function CheckoutItem({ className, data }) {
  var isCanceled = data.Canceled

  // Remain Unix
  var maxTime = data.MaxTimeConfirmation
  if (data.MaxTimeProcess != null) {
    maxTime = data.MaxTimeProcess
  }
  console.log(data.MaxTimeConfirmation, data.MaxTimeProcess, maxTime)

  return (
    <div className={className}>
      <div className='flex flex-col items-center justify-center w-full'>
        <div className='flex flex-row items-center px-12 w-full'>
          <p className={'font-bold ' + (isCanceled && "text-center")}>{data.Name}</p>
          {(!data.Delivered && !data.Canceled) && <TimeBadge targetTime={maxTime} />}
        </div>
        <div className='mt-3'>
          {
            !isCanceled ?
              <ul className="steps w-full">
                <li className={"step " + (data.Confirmation && "step-primary")}>Waiting for Confirmation</li>
                <li className={"step " + (data.Processed && "step-primary")}>On Processed</li>
                <li className={"step " + (data.Delivered && "step-primary")}>Delivered</li>
              </ul>
              :
              <div className='flex flex-col justify-center items-center'>
                <BsXCircle className='w-[40px] h-[40px]' color='#dc2626' />
                <p className="text-center text-red-500 font-semibold">Checkout Canceled</p>
              </div>
          }
        </div>
        <button class="btn btn-secondary btn-sm mt-8 max-w-xs" onClick={() => deleteChcekout(data.ID)}>Delete</button>
        <div className="divider"></div>
      </div>
    </div>
  )
}

async function deleteChcekout(id) {
  let res = await axios.delete(`http://localhost:3001/checkout/status/${id}`)
  if (res.status != 200) {
    alert("Cannot delete id " + id)
  }
}


export default CheckoutItem