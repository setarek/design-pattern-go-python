from abc import ABC, abstractmethod

# product
class Report:
    def __init__(self):
        self.title = None
        self.header = None
        self.content = None
        self.footer = None

    def display(self):
        print("Report:")
        print(f" Title: {self.title}")
        print(f" Header: {self.header}")
        print(f" Content: {self.content}")
        print(f" Footer: {self.footer}")

# builder interface
class ReportBuilder(ABC):
    
    @abstractmethod
    def set_header(self, header):
        pass

    @abstractmethod
    def set_title(self, tile):
        pass

    @abstractmethod
    def set_content(self, content):
        pass

    @abstractmethod
    def set_footer(self, footer):
        pass

    @abstractmethod
    def build(self):
        pass

# concerete builders

class PDFReportBuilder(ReportBuilder):
    def __init__(self):
        self.report = Report()

    def set_header(self, header):
        self.report.header = header
        return self
    
    def set_content(self, content):
        self.report.content = content
        return self

    def set_title(self, title):
        self.report.title = title
        return self

    def set_footer(self, footer):
        self.report.footer = footer
        return self

    def build(self) -> Report:
        return self.report
    
class CSVBuilder(ReportBuilder):
    def __init__(self):
        self.report = Report()
    def set_header(self, header):
        self.report.header = header
    
    def set_content(self, content):
        self.report.content = content
        return self

    def set_title(self, title):
        self.report.title = title
        return self

    def set_footer(self, footer):
        self.report.footer = footer
        return self

    def build(self) -> Report:
        return self.report

if __name__ == "__main__":
    
    pdf_builder = PDFReportBuilder()
    pdf_report = (
        pdf_builder
        .set_title("learning desing pattern")
        .set_header("Builder pattern")
        .set_content("lorem ipsum...")
        .set_footer("desing pattern cookbook")
        .build()
    )
    pdf_report.display()
    