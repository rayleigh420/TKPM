import { useContext } from "react";
import { useLocation, Navigate, Outlet } from "react-router-dom";
import AuthContext from "../../context/AuthProvider";

const RequireAuth = ({ role }) => {
    const { auth } = useContext(AuthContext);
    const location = useLocation();

    // console.log(admin, auth.admin)

    // if (auth) {
    //     auth.admin = Boolean(auth.admin)
    // }

    return (
        auth?.role == role
            ? <Outlet />
            : <Navigate to='/' state={{ from: location }} replace />
    );
}

export default RequireAuth;