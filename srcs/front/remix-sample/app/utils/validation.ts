// app/utils/validation.ts

export function validateUsername(username: string): string | undefined {
    if (username.length < 3) {
        return "ユーザー名は3文字以上で入力してください";
    } else if (username.length > 30) {
        return "ユーザー名は30文字以下で入力してください";
    }
    const usernameRegex = /^[a-zA-Z0-9_]+$/;
    if (!usernameRegex.test(username)) {
        return "有効なusernameを入力せよ"
    }
}

export function validateEmail(email: string): string | undefined {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        return "有効なメールアドレスを入力してください";
    }
}

export function validatePassword(password: string): string | undefined {
    if (password.length < 8) {
        return "パスワードは8文字以上で入力してください";
    } else if (password.length > 30) {
        return "パスワードは30文字以下で入力してください";
    }
}

export function validateRepassword(password: string, repassword: string): string | undefined {
    if (password !== repassword) {
        return "パスワードが一致しません";
    }
}


/**
 * 名前のバリデーション
 * - 空でないこと
 * - 文字種が適切であること（ひらがな、カタカナ、漢字）
 */
export function validateName(name: string): string | undefined {
    if (!name || name.trim().length === 0) {
        return "名前を入力してください";
    }

    // ひらがな、カタカナ、漢字のみを許可
    // const nameRegex = /^[ぁ-んァ-ン一-龯々ー]+$/;
    // if (!nameRegex.test(name)) {
    //     return "名前は日本語で入力してください";
    // }

    if (name.length > 20) {
        return "名前は20文字以内で入力してください";
    }
}

/**
 * 生年月日のバリデーション
 * - 必須
 * - 18歳以上
 * - 未来の日付でないこと
 */
export function validateBirthdate(birthdate: string): string | undefined {
    if (!birthdate) {
        return "生年月日を入力してください";
    }

    const birthdateObj = new Date(birthdate);
    const today = new Date();

    // 不正な日付
    if (isNaN(birthdateObj.getTime())) {
        return "正しい生年月日を入力してください";
    }

    // 未来の日付チェック
    if (birthdateObj > today) {
        return "生年月日が未来の日付になっています";
    }

    // 18歳以上チェック
    const age = today.getFullYear() - birthdateObj.getFullYear();
    const monthDiff = today.getMonth() - birthdateObj.getMonth();
    if (age < 18 || (age === 18 && monthDiff < 0)) {
        return "18歳未満の方は登録できません";
    }

    // 120歳以上チェック
    if (age > 120) {
        return "正しい生年月日を入力してください";
    }
}

/**
 * 性別のバリデーション
 * - 必須
 * - male または female のみ
 */
export function validateGender(gender: string): string | undefined {
    if (!gender) {
        return "性別を選択してください";
    }

    if (!["male", "female"].includes(gender)) {
        return "正しい性別を選択してください";
    }
}

/**
 * 希望する性別のバリデーション
 * - 必須
 * - male、female、male/female のいずれか
 */
export function validateSexuality(sexuality: string): string | undefined {
    if (!sexuality) {
        return "お相手の性別を選択してください";
    }

    if (!["male", "female", "male/female"].includes(sexuality)) {
        return "正しいお相手の性別を選択してください";
    }
}

/**
 * 地域のバリデーション
 * - 必須
 * - 有効な地域であること
 */
export const VALID_AREAS = [
    "北海道",
    "東京",
    "神奈川",
    "埼玉",
    "千葉",
    "大阪",
    "京都",
    "兵庫",
    "愛知",
    "福岡",
    "その他",
] as const;

export type Area = typeof VALID_AREAS[number];

export function validateArea(area: string): string | undefined {
    if (!area) {
        return "お住まいの地域を選択してください";
    }

    const validAreas = [
        'Hokkaido', 'Aomori', 'Iwate', 'Miyagi', 'Akita', 'Yamagata', 'Fukushima',
        'Ibaraki', 'Tochigi', 'Gunma', 'Saitama', 'Chiba', 'Tokyo', 'Kanagawa',
        'Niigata', 'Toyama', 'Ishikawa', 'Fukui', 'Yamanashi', 'Nagano',
        'Gifu', 'Shizuoka', 'Aichi', 'Mie',
        'Shiga', 'Kyoto', 'Osaka', 'Hyogo', 'Nara', 'Wakayama',
        'Tottori', 'Shimane', 'Okayama', 'Hiroshima', 'Yamaguchi',
        'Tokushima', 'Kagawa', 'Ehime', 'Kochi',
        'Fukuoka', 'Saga', 'Nagasaki', 'Kumamoto', 'Oita', 'Miyazaki', 'Kagoshima', 'Okinawa'
    ];
    if (!area || !validAreas.includes(area)) {
        return "都道府県を選択してください";
    }
}

/**
 * 自己紹介のバリデーション
 * - 必須
 * - 10文字以上1000文字以内
 * - 禁止ワードを含まないこと
 */
export function validateSelfIntro(selfIntro: string): string | undefined {
    if (!selfIntro) {
        return "自己紹介を入力してください";
    }

    const trimmedIntro = selfIntro.trim();

    if (trimmedIntro.length < 10) {
        return "自己紹介は10文字以上で入力してください";
    }

    if (trimmedIntro.length > 1000) {
        return "自己紹介は1000文字以内で入力してください";
    }

    // 禁止ワードチェック（必要に応じて追加）
    const prohibitedWords = ["暴力", "犯罪", "売春", "薬物"];
    const containsProhibitedWord = prohibitedWords.some(word =>
        trimmedIntro.includes(word)
    );

    if (containsProhibitedWord) {
        return "不適切な内容が含まれています";
    }
}

/**
 * 画像のバリデーション
 * - 必須
 * - ファイルサイズ制限（5MB以下）
 * - 許可された形式（JPEG, PNG, GIF）
 */
export function validateImage(file: File): string | undefined {
    if (!file) {
        return "プロフィール画像を選択してください";
    }

    // ファイルサイズチェック（5MB）
    const maxSize = 5 * 1024 * 1024;
    if (file.size > maxSize) {
        return "画像サイズは5MB以下にしてください";
    }

    // ファイル形式チェック
    const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
    if (!allowedTypes.includes(file.type)) {
        return "JPEG、PNG、GIF形式の画像のみアップロード可能です";
    }
}

/**
 * フォーム全体のバリデーション
 */
export interface FormData {
    lastname: string;
    firstname: string;
    birthdate: string;
    gender: string;
    sexuality: string;
    area: string;
    selfIntro: string;
    image: File;
}

export interface ValidationErrors {
    lastname?: string;
    firstname?: string;
    birthdate?: string;
    gender?: string;
    sexuality?: string;
    area?: string;
    selfIntro?: string;
    image?: string;
}

export function validateForm(data: FormData): ValidationErrors {
    const errors: ValidationErrors = {};

    const lastnameError = validateName(data.lastname);
    if (lastnameError) errors.lastname = lastnameError;

    const firstnameError = validateName(data.firstname);
    if (firstnameError) errors.firstname = firstnameError;

    const birthdateError = validateBirthdate(data.birthdate);
    if (birthdateError) errors.birthdate = birthdateError;

    const genderError = validateGender(data.gender);
    if (genderError) errors.gender = genderError;

    const sexualityError = validateSexuality(data.sexuality);
    if (sexualityError) errors.sexuality = sexualityError;

    const areaError = validateArea(data.area);
    if (areaError) errors.area = areaError;

    const selfIntroError = validateSelfIntro(data.selfIntro);
    if (selfIntroError) errors.selfIntro = selfIntroError;

    const imageError = validateImage(data.image);
    if (imageError) errors.image = imageError;

    return errors;
}