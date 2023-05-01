import { useContext } from "react";
import { useLocation, Navigate, Outlet } from "react-router-dom";
import AuthContext from "../../context/AuthProvider";

const RequireAuth = ({ allowedRoles }) => {
    const { auth } = useContext(AuthContext);
    const location = useLocation();

    console.log(allowedRoles)

    // console.log(admin, auth.admin)

    // if (auth) {
    //     auth.admin = Boolean(auth.admin)
    // }

    return (
        allowedRoles.includes(auth?.role)
            ? <Outlet />
            : <Navigate to='/' state={{ from: location }} replace />
    );
}

export default RequireAuth;