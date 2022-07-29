import React, { useState } from 'react'
import axios from 'axios'

function InputCard({className}) {
  const [Name, setName] = useState("")

  async function postCheckout() {
    let res = await axios.post("http://localhost:3001/checkout", {
      "name" : Name
    })
    if (res.status != 200) {
      alert("Gagal membuat checkout")
    }
    setName("")
  }

  return (
    <div className={className}>
      <div className="card bg-base-200 shadow-xl">
        <div className="card-body">
          <h2 className="card-title">Masukkan Nama Checkout</h2>
          <input type="text" placeholder="Nama checkout" value={Name} onChange={v => {setName(v.target.value)}} className="input input-bordered input-primary w-full" />
          <button onClick={postCheckout} className="btn btn-primary">Kirim</button>
        </div>
      </div>      
    </div>
  )
}

export default InputCard