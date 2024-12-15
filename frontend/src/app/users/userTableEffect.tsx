import { User, GetPermissionName } from "@/model/user";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';

type UserTableEffectProps = {
    users: User[];
}

const UserTableEffect: React.FC<UserTableEffectProps> = ({ users }) => {
    console.log(users);

    return(
        <>
            <TableContainer>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>ユーザID</TableCell>
                            <TableCell>メールアドレス</TableCell>
                            <TableCell>権限</TableCell>
                            <TableCell>職業</TableCell>
                            <TableCell></TableCell>
                         </TableRow>
                    </TableHead>
                    <TableBody>
                        {users.map((user: User) => (
                            <TableRow>
                                <TableCell>{user.user_id}</TableCell>
                                <TableCell>{user.email}</TableCell>
                                <TableCell>{GetPermissionName(user.role)}</TableCell>
                                <TableCell>{user.job}</TableCell>
                                {/* <TableCell>
                                    <form action={DeleteUser}>
                                        <input type="hidden" name="id" value={user.user_id}/>
                                        <button type="submit">削除</button>
                                    </form> */}
                                {/* </TableCell> */}
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}

export default UserTableEffect;