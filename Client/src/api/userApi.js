import axios from "../utils/axios"

export const getAllUser = async (token) => {
    console.log(token)
    const result = await axios.get('/admin/users', {
        token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcjEiLCJlbWFpbCI6InVzZXIxQGdtYWlsLmNvbSIsInJvbGUiOiJhZG1pbiIsInVzZXJfaWQiOiJVMDEiLCJhdmF0YXIiOiJodHRwczovL2Zhc3RseS5waWNzdW0ucGhvdG9zL2lkLzY3Ni8yMDAvMjAwLmpwZz9obWFjPWhnZU1RRUlLNE1uMjdRMm9MUldqWG8xcmd4d1RiazFDbkpFOTU0aF9IeU0iLCJleHAiOjE2ODM3MjczMjl9.I2LcAstsPd5Pw8ODhznmjkTS_JZeqRulJVyR-I3jf1I"
    })
    console.log(result.data)
    return result.data
}