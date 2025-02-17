USE [sqlboiler_test]
GO
/****** Object:  Table [dbo].[TestTable]    Script Date: 8/6/2024 1:46:58 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TestTable](
	[Id] [uniqueidentifier] NOT NULL,
	[Date] [datetime] NOT NULL,
 CONSTRAINT [PK_TestTable] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'8de4dd9b-46d5-4141-9a72-000d96698e74', CAST(N'2019-05-03T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'229e74fc-31b3-42f9-aba3-0028351aca30', CAST(N'2022-02-18T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'5ab0412a-2b2f-430f-8830-002a42125148', CAST(N'2023-09-15T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'84ae9595-8e76-4d8c-a08c-00434be75672', CAST(N'2019-07-12T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'1fd44e8f-0494-42c6-b21a-0072e1b8c2f2', CAST(N'2014-07-08T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'192452f8-93c2-4a20-a52b-0093741a9e45', CAST(N'2015-03-27T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'3f30584c-31e6-4227-96b1-00ce1acb47ad', CAST(N'2019-05-03T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'3066d9dd-2e98-4505-94fa-01062331ed1a', CAST(N'2021-07-23T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'e4e4d205-8721-46c4-81b9-017c9a5adcae', CAST(N'2023-02-03T00:00:00.000' AS DateTime))
INSERT [dbo].[TestTable] ([Id], [Date]) VALUES (N'c0ba42c9-480c-4236-b217-01910e51b290', CAST(N'2020-07-03T00:00:00.000' AS DateTime))
GO
