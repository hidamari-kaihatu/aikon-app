import React from "react";

export default function Today() {
    const today = new Date();
    const year = today.getFullYear()
    const month = today.getMonth() + 1
    const day = today.getDate()
    const week = today.getDay()
    const weekItems = ["日", "月", "火", "水", "木", "金", "土"]
    const dayOfWeek = weekItems[week]
    return (
        <div suppressHydrationWarning>{year}年{month}月{day}日（{dayOfWeek}）</div>
    );
}