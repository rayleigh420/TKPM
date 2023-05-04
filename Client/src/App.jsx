import './App.css'
import { Route, Routes } from 'react-router-dom'
import Layout from './pages/Layout'
import Home from './pages/home'
import Book from './pages/Book'
import UserLayout from './pages/UserLayout'
import ProfileForm from './components/forms/ProfileForm'
import ListRentedBook from './components/user/ListRentedBook'
import ManageUser from './components/admin/manages/ManageUser'
import ManageBook from './components/admin/manages/ManageBook'
import BookedBook from './components/status/BookedBook'
import BorrowedBook from './components/status/BorrowedBook'
import PaidBook from './components/status/PaidBook'
import GenericBook from './pages/GenericBook'
import RequireAuth from './components/auth/RequireAuth'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route path='/' element={<Home />} />
          <Route path='book/:id' element={<Book />} />
          <Route path='user' element={<UserLayout />} >

            <Route element={<RequireAuth allowedRoles={["user", "admin"]} />}>
              <Route path='profile' element={<ProfileForm />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["user"]} />}>
              <Route path='listrented' element={< ListRentedBook />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["admin"]} />}>
              <Route path='manage/viewer' element={<ManageUser />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["admin"]} />}>
              <Route path='manage/book' element={<ManageBook />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["admin"]} />}>
              <Route path='status/booked' element={<BookedBook />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["admin"]} />}>
              <Route path='status/borrowed' element={<BorrowedBook />} />
            </Route>

            <Route element={<RequireAuth allowedRoles={["admin"]} />}>
              <Route path='status/paid' element={<PaidBook />} />
            </Route>

          </Route>
          <Route path="/:generic" element={<GenericBook />} />
        </Route>
      </Routes>
    </>
  )
}

export default App
