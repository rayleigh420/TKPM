import './App.css'
import { Route, Routes } from 'react-router-dom'
import Layout from './pages/Layout'
import Home from './pages/home'
import Book from './pages/Book'
import ListRentedBook from './pages/ListRentedBook'
import UserLayout from './pages/UserLayout'
import ProfileForm from './components/ProfileForm'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route path='home' element={<Home />} />
          <Route path='book' element={<Book />} />
          <Route path='user' element={<UserLayout />} >
            <Route path='profile' element={<ProfileForm />} />
            <Route path='listrented' element={< ListRentedBook />} />
          </Route>
        </Route>
      </Routes>
    </>
  )
}

export default App
