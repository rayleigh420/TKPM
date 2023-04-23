import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Navigation from './components/navigation'
import Footer from './components/Footer'

const data = [

  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  },
  {
    src: 'https://i.truyenvua.com/ebook/190x247/boi-hoi-gia_1663135088.jpg?gt=hdfgdfg&mobile=2',
    name: 'Boi hoi gia'
  }
]

function App() {

  return (
    <>
      <Navigation />
      <div className="px-[135px] flex gap-5 flex-wrap mt-[30px]">
        {
          data.map(item => (
            <div className='w-[158px] h-[270px] rounded-[7px] cursor-pointer'>
              <img src={item.src} className='w-full h-[230px]' />
              <p className='text-[#e0e0e0] text-[15.2px] font-semibold leaiding-[18.24px] tracking-[-0.304px] mt-1'>{item.name}</p>
            </div>
          ))
        }
      </div>
      <Footer />
    </>
  )
}

export default App
