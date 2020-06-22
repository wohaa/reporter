/*
   Copyright 2016 Vastech SA (PTY) LTD

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package report

const defaultRowTemplate = `

%use square brackets as golang text templating delimiters
\documentclass{article}

\usepackage{helvet}
\renewcommand{\familydefault}{\sfdefault}
\usepackage{graphicx,fancyhdr,xcolor,xifthen}
\usepackage{hyperref}
\hypersetup{
    colorlinks = false,
    linkbordercolor = {white}
}
\fancyhead[R]{}
\fancyhead[L]{\includegraphics[width=4cm]{/path/to/your/logo.png}}
%\fancyfoot[R]{\thepage}
\renewcommand{\headrulewidth}{0pt}
\usepackage[margin=1in]{geometry}
\graphicspath{ {images/} }
\begin{document}
\title{[[.Title]] [[if .VariableValues]] \\ \large [[.VariableValues]] [[end]] [[if .Description]] \\ \small [[.Description]] [[end]]}
\author{Row Template Report}
\maketitle
{\centering
	\vspace{3cm}
	\par
	Evaluation period:
	\par
	from \date{[[.FromFormatted]] \\ to [[.ToFormatted]]}
	\vspace{3cm}
	\par
	\vspace{8cm}
	\par
	\small Author
	\par
	\thispagestyle{fancy}
	%Space for general informations
}

\pagebreak
\tableofcontents
\pagebreak

[[range .Panels]][[if .IsSingleStat]]\begin{minipage}{0.3\textwidth}
\includegraphics[width=\textwidth]{image[[.Id]]}
\thispagestyle{fancy}
\end{minipage}
[[else]]\par
\ifthenelse{\equal{[[.Section]]}{}}{}{\pagebreak\section{[[.Section]]}}
%\section{[[.Section]]}
\vspace{0.5cm}
\subsection{[[.Title]]}
%\subsection{[[.Title]]}
\includegraphics[width=\textwidth]{image[[.Id]]}
\thispagestyle{fancy}
\par
%Space for further informations
\vspace{0.5cm}
[[end]][[end]]

\end{document}
`
