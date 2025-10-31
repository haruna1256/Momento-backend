# Momento Backend API

個人開発アプリ「Momento」用のバックエンドAPIです。  
Goで実装し、最終的にKubernetes上で稼働させることを目標としています。

## 技術構成
- Language: Go
- Framework: net/http（→ Ginへ拡張予定）
- Auth: 自作JWT認証（予定）
- DB: SQLite or PostgreSQL（後で検討）
- Deployment: Docker + kind (Kubernetes)

## 機能予定
- ユーザー登録・認証・認可
- 写真・アルバムデータのCRUD
- Before/After履歴の取得API

## 開発メモ
- フロントエンド（SwiftUI）とは別リポジトリ
- 現在はAPI実験フェーズ
