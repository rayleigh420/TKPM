import { Outlet } from "react-router-dom"
import Footer from "../components/layout/Footer"
import Form from "../components/Form"

const Layout = () => {
    return (
        <>
            <Form />
            <div className="pt-[65px]">
                <Outlet />
            </div>
            <Footer />
        </>
    )
}

export default Layout