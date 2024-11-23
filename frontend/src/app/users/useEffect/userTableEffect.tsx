import { User, GetPermissionName } from "@/model/user";
import { type FC } from 'react'
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Link from 'next/link'
import { DeleteUser } from "./fetcher";

type UserTableEffectProps = {
    users: User[];
}

const UserTableEffect: FC<UserTableEffectProps> = ({ users }) => {

    return(
        <>
            <TableContainer>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>名前</TableCell>
                            <TableCell>メールアドレス</TableCell>
                            <TableCell>権限</TableCell>
                            <TableCell>職業</TableCell>
                            <TableCell></TableCell>
                         </TableRow>
                    </TableHead>
                    <TableBody>
                        {users.map((user: User) => (
                            <TableRow key={user.id}>
                                <TableCell>{user.name}</TableCell>
                                <TableCell>{user.email}</TableCell>
                                <TableCell>{GetPermissionName(user.role)}</TableCell>
                                <TableCell>{user.job}</TableCell>
                                <TableCell>
                                    <form action={DeleteUser}>
                                        <input type="hidden" name="id" value={user.id}/>
                                        <button type="submit">削除</button>
                                    </form>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}

export default UserTableEffect;