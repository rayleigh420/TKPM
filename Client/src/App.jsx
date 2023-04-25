import './App.css'
import { Route, Routes } from 'react-router-dom'
import Layout from './pages/Layout'
import Home from './pages/home'
import Book from './pages/Book'
import UserLayout from './pages/UserLayout'
import ProfileForm from './components/ProfileForm'
import ListRentedBook from './components/ListRentedBook'
import ManageUser from './components/ManageUser'
import ManageBook from './components/ManageBook'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route path='/' element={<Home />} />
          <Route path='book' element={<Book />} />
          <Route path='user' element={<UserLayout />} >
            <Route path='profile' element={<ProfileForm />} />
            <Route path='listrented' element={< ListRentedBook />} />
            <Route path='manage/viewer' element={<ManageUser />} />
            <Route path='manage/book' element={<ManageBook />} />
          </Route>
        </Route>
      </Routes>
    </>
  )
}

export default App
