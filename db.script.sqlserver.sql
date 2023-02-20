
CREATE TABLE [dbo].[person](
	[row] [int] IDENTITY(1,1) NOT NULL,
	[id] [char](10) NOT NULL,
	[name] [nvarchar](10) NOT NULL,
	[age] [int] NOT NULL,
	[birthday] [datetime] NOT NULL,
	[remark] [nvarchar](100) NULL,
 CONSTRAINT [pk__person] PRIMARY KEY CLUSTERED
(
	[row] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

CREATE UNIQUE NONCLUSTERED INDEX [uni__person__id] ON [dbo].[person]
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO

CREATE TABLE [dbo].[address](
	[row] [int] IDENTITY(1,1) NOT NULL,
	[id] [char](10) NOT NULL,
	[text] [nvarchar](100) NOT NULL,
 CONSTRAINT [pk__address] PRIMARY KEY CLUSTERED
(
	[row] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[address]  WITH CHECK ADD CONSTRAINT [fk__address__id] FOREIGN KEY([id])
REFERENCES [dbo].[person] ([id])
ON UPDATE CASCADE
ON DELETE CASCADE
GO

ALTER TABLE [dbo].[address] CHECK CONSTRAINT [fk__address__id]
GO

-- Data
INSERT INTO [dbo].[person]
([id],[name],[age],[birthday],[remark])
VALUES
('A123456789','AAA',18,GETDATE(),''),
('B123456789','BBB',28,GETDATE(),'')
GO

INSERT INTO [dbo].[address]
([id],[text])
VALUES
('A123456789','A1'),
('A123456789','A2'),
('B123456789','B1'),
('B123456789','B2')
GO