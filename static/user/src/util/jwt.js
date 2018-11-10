import Cookies from "js-cookie";

const GetJWTTokenString = () =>{
    return Cookies.get("yougame_token")
};
export default GetJWTTokenString;