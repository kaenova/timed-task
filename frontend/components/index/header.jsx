import React from 'react'

function MainHeader({className}) {
  return (
    <div className={className}>
      <div className="card bg-base-200 shadow-xl">
        <div className="card-body">
          <h2 className="card-title">Halo selamat datang!</h2>
          <p>Ini merupakan percobaan dari sistem checkout yang dibuat dengan menggunakan Delayed Message RabbitMQ</p>
          <p>Created by <a href="https://kaenova.my.id/" className='underline text-primary'>Kaenova Mahendra Auditama</a></p>
        </div>
      </div>      
    </div>
  )
}

export default MainHeader