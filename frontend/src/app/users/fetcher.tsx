"use server"

import { User } from "@/model/user";
import { cookies } from "next/headers";

export async function GetUsers() {
    const cookieStore = await cookies();
    const cookie = cookieStore.get('Authorization')?.value;
    const headers = new Headers();
    if (cookie) {
        headers.append('cookie', cookie);
    }
    const res = await fetch(`http://backend:1323/api/users`, {
        method: 'GET',
        headers,
    });
    const data = await res.json();
    console.log("server")
    // const resCookie = res.headers.get('set-cookie');
    // if (resCookie) {
    //     const cookieStore = await cookies();
    //     cookieStore.set({name: 'Authorization', value: resCookie});
    // }
    const users: User[] = data.users;
    return users;
}

// export async function CreateUser(
//     data: FormData,
// ) {
//     const users: User[] = await GetUsers();
//     const { name, email, role, job } = Object.fromEntries(data.entries());

//     await fetch(`http://jsonserver:3001/users`, {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//         },
//         body: JSON.stringify({ name, email, role, job}),
//     });
//     revalidatePath('/users/useEffect')
// }

// export async function DeleteUser(
//     data: FormData,
// ) {
//     const id = data.get('user_id');
//     await fetch(`http://jsonserver:3001/users/${id}`, {
//         method: 'DELETE',
//     });
//     revalidatePath('/users/useEffect')
// }