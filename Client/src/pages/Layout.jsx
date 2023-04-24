import { Outlet } from "react-router-dom"
import Footer from "../components/Footer"
import Form from "../components/Form"

const Layout = () => {
    return (
        <>
            <Form />
            <Outlet />
            <Footer />
        </>
    )
}

export default Layout