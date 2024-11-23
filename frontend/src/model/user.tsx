"use server";

export interface User {
    id: number;
    name: string;
    email: string;
    role: string;
    job: string;
}

const permissionMap: Record<string, string> = {
    '01': '管理者',
    '02': '一般',
    '03': '閲覧者',
    '04': 'ゲスト',
};

export async function GetPermissionName (permissionId: string): Promise<string>  {
    return permissionMap[permissionId] || '不明'; // マッピングにない場合は「不明」
  };