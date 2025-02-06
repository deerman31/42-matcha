export const saveToken = (token: string) => {
  // HttpOnly cookieが最も安全ですが、バックエンドで設定する必要があります
  // フロントエンドではsessionStorageを使用するのが次善の策です
  sessionStorage.setItem("access_token", token);
};

export const getToken = (): string | null => {
  return sessionStorage.getItem("access_token");
};

export const removeToken = () => {
  sessionStorage.removeItem("access_token");
};
