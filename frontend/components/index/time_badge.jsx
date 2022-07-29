import React, { useEffect, useState } from 'react'
import timestamp from 'unix-timestamp';

function TimeBadge({ targetTime }) {
  console.log(targetTime)
  // Convert to milisecond
  targetTime = targetTime * 1000

  // Timeleft
  let timeLeft = targetTime - new Date()
  if (timeLeft < 0) {
    timeLeft = 0
  }

  const [TimeLeft, setTimeLeft] = useState(timeLeft)
  useEffect(() => {
    let inter = setInterval(() => {
      setTimeLeft(TimeLeft)
      TimeLeft = TimeLeft - 1000
      if (TimeLeft < 0) {
        TimeLeft = 0
      }
    }, 1000)
    return () => {
      clearInterval(inter)
    }
  }, [])


  return (
    <div className="badge w-[160px]">Canceled in {dateToString(TimeLeft)}</div>
  )
}

function dateToString(milisecond) {
  let date = new Date(milisecond)
  return date.toISOString().substring(11, 19)
}

export default TimeBadge