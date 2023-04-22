import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Navigation from './components/navigation'
import Footer from './components/Footer'

function App() {

  return (
    <>
      <Navigation />
      <div className="artboard phone-6 px-[140px]">
        <div className='w-[158px] h-[270px] rounded-[10px] cursor-pointer'>
          <img src="https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2" className='w-full h-[230px] overflow-hidden'></img>
          <p className='text-[#e0e0e0] text-[15.2px] font-semibold leaiding-[18.24px] tracking-[-0.304px] mt-1'>Boi hoi gia</p>
        </div>
      </div>
      <main className='px-[140px]'>Hello</main>
      <Footer />
    </>
  )
}

export default App
