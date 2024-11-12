"use server";

import { User } from "@/model/user";
import { revalidatePath } from "next/cache";

export async function GetUsers() {
    const res = await fetch(`http://jsonserver:3001/users`);
    return res.json();
}

export async function CreateUser(
    data: FormData,
) {
    const users: User[] = await GetUsers();
    const maxId = users.length > 0 ? Math.max(...users.map(user => user.id)) : 0;
    const id = maxId + 1;
    const { name, email, role, job } = Object.fromEntries(data.entries());

    await fetch(`http://jsonserver:3001/users`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id, name, email, role, job}),
    });
    revalidatePath('/users/useEffect')
}

export async function DeleteUser(
    data: FormData,
) {
    const id = data.get('id');
    await fetch(`http://jsonserver:3001/users/${id}`, {
        method: 'DELETE',
    });
    revalidatePath('/users/useEffect')
}