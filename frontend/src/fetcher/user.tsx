"use server";
import { User } from "@/model/user";

export async function GetUsers(
    _prevState: User[],
    formData: FormData,
) {
    const query = formData.get("query") as string;
    const res = await fetch(`http://jsonserver:3001/users?q=${query}`);
    const users: User[] = await res.json();
    return users;
}