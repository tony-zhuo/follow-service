# model
為資料訪問層
1.	Model 層主要負責定義系統的業務邏輯和核心規則。這些模型（Entities）是應用程序的核心對象，反映了應用程序在最高層次的商業邏輯和規則。這層應該是最穩定和不受外界變化影響的部分。
2.  這層通常包含實體（Entities）、值對象（Value Objects）、和聚合（Aggregates）等，這些對象封裝了業務規則，並且不依賴於應用程序的其他層或技術細節。