// src/types/api.ts として型定義を分離

export interface Response {
  message: string;
}

export interface LoginResponse {
  is_preparation: boolean;
  access_token: string;
}

export interface ErrorResponse {
  error: string;
}

export interface MyInfo {
  username: string;
  email: string;
  lastname: string;
  firstname: string;
  birthdate: string;
  gender: string;
  sexuality: string;
  area: string;
  self_intro: string;
  tags: string[];
  is_gps: boolean;
  latitude: number;
  longitude: number;
  fame_rating: number;
}

export interface MyProfileResponse {
  my_info: MyInfo;
}

export interface UserInfo {
  username: string;
  age: number;
  distance_km: number;
  common_tag_count: number;
  fame_rating: number;
  image_path: string;
}
export interface BrowseResponse {
  user_infos: UserInfo[];
}

export interface BrowseUserImageResponse {
  image: string;
}

export interface AllMyImageResponse {
  all_image: string[];
}

export interface OtherProfile {
  username: string;
  age: number;
  gender: string;
  sexuality: string;
  area: string;
  self_intro: string;
  tags: string[];
  distance: number;
  fame_rating: number;
}

export interface OtherProfileResponse {
  other_profile: OtherProfile;
  error?: string;
}

export interface AllOtherImageResponse {
  all_image: string[];
}
